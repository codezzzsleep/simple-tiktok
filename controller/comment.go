package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
// CommentAction并不起作用，只是检查token是否合法
// CommentAction 发表评论、删除评论
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	//判断用户是否登录
	if token == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "You haven't logged in yet",
		})
	}
	action := c.Query("action_type")
	actionType, err := strconv.Atoi(action)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	if actionType == 1 {
		PostComment(c)
	} else {
		DeleteComment(c)
	}

}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	//获取需要删除的评论Id
	commentIdStr := c.Query("comment_id")
	commentId, _ := strconv.ParseInt(commentIdStr, 10, 64)

	videoIdStr := c.Query("video_id")
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)

	//（1）直接在数据库中删除（2）comment中设置一个deleted 列，用bool表示是否删除
	//目前实现的是第一种
	//开启数据库事务，在comments中添加记录，在videos中更改评论数目
	DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id=?", commentId).Delete(&Comment{}).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		tx.Model(&Video{}).Where("id=?", videoId).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1))
		// 返回 nil 提交事务
		return nil
	})
	c.JSON(http.StatusOK, Response{
		StatusCode: 1,
		StatusMsg:  "comment deleted successfully",
	})
}

// PostComment 发表评论
func PostComment(c *gin.Context) {
	username := c.Query("token")
	//根据用户名查找用户
	var user User
	DB.Where("name=?", username).First(&user)
	//读取评论内容
	context := c.Query("comment_text")
	//用户未输入评论内容
	if context == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "请输入评论内容",
		})
		return
	}
	videoIdStr := c.Query("video_id")
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	//日期MM-DD
	timeFormat := time.Now().Format("01-02")
	comment := Comment{
		User:       user,
		VideoId:    videoId,
		UserId:     user.Id,
		Content:    context,
		CreateDate: timeFormat,
	}
	//开启数据库事务，在comments中添加记录，在videos中更改评论数目
	DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&comment).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		tx.Model(&Video{}).Where("id=?", videoId).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1))
		// 返回 nil 提交事务
		return nil
	})

	//文档中标明不需要拉取评论列表，数据库中的自增id无法获取
	//目前默认每次处理一条comment，所以数组只存入一条评论数据
	commentList := []Comment{comment}

	c.JSON(http.StatusOK, CommentListResponse{
		Response: Response{StatusCode: 0,
			StatusMsg: "comment posted successfully"},
		CommentList: commentList,
	})
}

// CommentList 获取评论列表
func CommentList(c *gin.Context) {
	token := c.Query("token")
	//判断用户是否登录
	if token == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "You haven't logged in yet",
		})
	}

	videoIdStr := c.Query("video_id")
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)

	var commentList []Comment
	DB.Preload("User").Where("video_id=?", videoId).Find(&commentList)

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0, StatusMsg: "success"},
		CommentList: commentList,
	})
}
