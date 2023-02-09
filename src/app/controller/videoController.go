package controller

import (
	"DouYin/src/app/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	NextTime  int64           `json:"next_time"`
	VideoList []service.Video `json:"video_list"`
}

type VideoListResponse struct {
	Response
	VideoList []service.Video `json:"video_list"`
}

func GetVideo() service.VideoServiceImpl {
	var userService service.UserServiceImpl
	var videoService service.VideoServiceImpl
	videoService.UserService = &userService
	return videoService
}

// 视频流接口
func Feed(ctx context.Context, c *app.RequestContext) { //接受视频处理
	inputTime := c.Query("latest_time")
	log.Printf("传入的时间" + inputTime)
	var lastTime time.Time
	if len(inputTime) != 0 {
		msec, _ := strconv.ParseInt(inputTime, 10, 64)
		lastTime = time.UnixMilli(msec)
	} else {
		lastTime = time.Now()
	}
	log.Printf("获取到时间戳%v", lastTime)
	userId, _ := strconv.ParseInt(c.GetString("userId"), 10, 64)
	log.Printf("获取到用户id:%v\n", userId)
	videoService := GetVideo()
	feed, nextTime, err := videoService.Feed(lastTime)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{StatusCode: 1, StatusMsg: "获取视频流失败"},
		})
		return
	}
	log.Println(nextTime)
	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{StatusCode: 0},
		NextTime: time.Now().UnixMilli(),
		//NextTime:  nextTime.UnixMilli(),
		VideoList: feed,
	})
}

// 上传视频
func PublishAction(ctx context.Context, c *app.RequestContext) {

}

// 获取用户上传列表
func GetPublishList(ctx context.Context, c *app.RequestContext) {
	strUserId := c.Query("user_id")
	userId, _ := strconv.ParseInt(strUserId, 10, 64)
	videoService := GetVideo()
	publishList, err := videoService.GetPublishVideoList(userId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Get PublishList failed"},
		})
	} else {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 1, StatusMsg: "Get PublishList suc"},
			VideoList: publishList,
		})
	}
}
