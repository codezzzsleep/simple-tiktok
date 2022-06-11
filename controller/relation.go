package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
// RelationAction函数并不起作用，只是检查返回的token是否合法
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	//判断用户是否登录
	if token == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "You haven't logged in yet",
		})
	}
	//赞请求的具体操作
	action := c.Query("action_type")
	actionType, err := strconv.Atoi(action)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//判断活动类型
	if actionType == 1 {
		Follow(c)
	} else {
		Unfollow(c)
	}
}

// FollowList all users have same follow list
//所有人都有相同的关注列表
func FollowList(c *gin.Context) {
	username := c.Query("token") // 用户名
	if username == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "You haven't logged in yet",
		})
	}
	var user User
	DB.Where("name=?", username).First(&user)
	var followList []Relation
	// 加载 UserB 即加载当前用户关注的用户
	DB.Preload("UserB").Where("user_a_id=?", user.Id).Find(&followList)
	// 复制用户关注列表信息
	followUserList := make([]User, len(followList))
	for i, f := range followList {
		followUserList[i] = f.UserB
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: followUserList,
	})
}

// FollowerList all users have same follower list
//所有人都有相同的粉丝列表
func FollowerList(c *gin.Context) {
	username := c.Query("token") // 用户名
	if username == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "You haven't logged in yet",
		})
	}
	var user User
	DB.Where("name=?", username).First(&user)
	var followerList []Relation
	// 加载 UserA 即加载当前用户的粉丝
	DB.Preload("UserA").Where("user_b_id=?", user.Id).Find(&followerList)
	// 复制用户粉丝列表信息
	followerUserList := make([]User, len(followerList))
	for i, f := range followerList {
		followerUserList[i] = f.UserA
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: followerUserList,
	})
}

// Follow 关注功能
func Follow(c *gin.Context) {
	username := c.Query("token")
	// 根据用户名查找要关注的用户
	var user User
	DB.Where("name=?", username).First(&user)
	userIdToStr := c.Query("to_user_id") // UserB
	userIdTo, _ := strconv.ParseInt(userIdToStr, 10, 64)
	relation := Relation{
		UserAId: user.Id,
		UserBId: userIdTo,
	}
	// 开启数据库事务，在 relations 中添加记录，在 users 中更改关注数
	DB.Transaction(func(tx *gorm.DB) error {
		// 创建关注关系
		if err := tx.Create(&relation).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		// 添加关注数、被关注数(粉丝数)
		tx.Model(&User{}).Where("id=?", user.Id).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1))
		tx.Model(&User{}).Where("id=?", userIdTo).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1))
		// 返回 nil 提交事务
		return nil
	})
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "Successfully follow",
	})
}

// Unfollow 取消关注
func Unfollow(c *gin.Context) {
	userIdToStr := c.Query("to_user_id") //UserB
	userIdTo, _ := strconv.ParseInt(userIdToStr, 10, 64)
	username := c.Query("token")
	// 根据用户名查找要取关的用户
	var user User
	DB.Where("name=?", username).First(&user)
	// 开启数据库事务，在 relations 中添加记录，在 users 中更改关注数
	DB.Transaction(func(tx *gorm.DB) error {
		// 删除关注关系
		if err := tx.Where("user_a_id=?", user.Id).Where("user_b_id=?", userIdTo).Delete(&Relation{}).Error; err != nil {
			//返回任何错误都会回滚事务
			return err
		}
		// 减少关注数、被关注数
		tx.Model(&User{}).Where("id=?", user.Id).UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1))
		tx.Model(&User{}).Where("id=?", userIdTo).UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1))
		//返回nil提交事务
		return nil
	})
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "Successfully unfollowed",
	})
}
