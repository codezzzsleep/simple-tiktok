package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
// Publish 函数检查 token 是否合法，然后将上传视频保存到 public 文件夹中
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	var user User
	if rowsAffected := DB.Where("name = ?", token).First(&user).RowsAffected; rowsAffected == 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist3"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	// 视频标题
	title := c.PostForm("title")
	filename := filepath.Base(data.Filename)
	playUrl := "/static/"
	//存视频url的前半部分,存的都是相对路径，在播放的时候把路径加载出来之后，前面添加本机IP
	coverUrl := "/static/image/"
	//存视频封面url的前半部分,存的都是相对路径，在播放的时候把路径加载出来之后，前面添加本机IP
	video1 := Video{
		Author:   user,
		UserId:   user.Id,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    title,
	}
	//用户自己还没搞，先暂时使用固定的用户user1.已修改为当前登录用户
	DB.Create(&video1) //创建数据
	finalName := fmt.Sprintf("%d_%d_%s", video1.Id, user1.Id, filename)
	//为了避免重复前面增加一个视频id，也可以搞一个随机数，减少重复几率
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})

	//接下来为视频寻找封面，直接选择10帧
	cmd := exec.Command(
		"ffmpeg", "-i", "./public/"+finalName,
		"-vf", "select=eq(n\\, 10)", "-frames", "1",
		"./public/image/"+finalName+".jpg",
	)
	//此处需安装ffmpeg,教程如下https://blog.csdn.net/m0_53574178/article/details/122565831
	//cmd.Stderr = os.Stderr // 输出错误信息
	if err := cmd.Run(); err != nil {
		log.Fatalln("视频封面截取失败")
	}
	// 更新数据库中的视频和封面链接
	video1.PlayUrl += finalName
	video1.CoverUrl += finalName + ".jpg"
	DB.Save(&video1)
}

// PublishList all users have same publish video list
//每个人都有相同的 发布视频列表
func PublishList(c *gin.Context) {
	fmt.Println("显示当前用户的所有发布视频s")
	fmt.Println(c.Query("user_id"))
	var videoList []Video //创建一个Video数组存起来
	DB.Preload("Author").Where("user_id = ?", c.Query("user_id")).Find(&videoList)
	for i := range videoList {
		fmt.Println(i)
		videoList[i].CoverUrl = "http://172.33.170.41:8080" + videoList[i].CoverUrl
		videoList[i].PlayUrl = "http://172.33.170.41:8080" + videoList[i].PlayUrl
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
