package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
// 现在的 VideoList 只能获取已经预先写好的 DemoVideos 变量

func Feed(c *gin.Context) {
	var videoList []Video //创建一个Video数组存起来
	DB.Preload("Author").Order("id desc").Limit(30).Find(&videoList)
	for i := range videoList {
		videoList[i].CoverUrl = "http://172.33.170.41:8080" + videoList[i].CoverUrl
		videoList[i].PlayUrl = "http://172.33.170.41:8080" + videoList[i].PlayUrl
	}
	//为url添加前缀
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}
