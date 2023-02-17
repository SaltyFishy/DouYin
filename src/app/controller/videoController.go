package controller

import (
	"DouYin/src/app/middleware/jwt"
	"DouYin/src/app/model"
	"DouYin/src/app/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"mime/multipart"
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
	var favoriteService service.FavoriteServiceImpl
	var commentService service.CommentServiceImpl
	videoService.UserService = &userService
	videoService.FavoriteService = &favoriteService
	videoService.CommentService = &commentService
	return videoService
}

// 视频流接口
func Feed(ctx context.Context, c *app.RequestContext) { //接受视频处
	var lastTime, nextTime time.Time
	var err error
	var userId int64
	var token interface{}
	var ok bool = false
	var feed []service.Video

	inputTime := c.Query("latest_time")
	log.Printf("传入的时间" + inputTime)
	if len(inputTime) != 0 {
		msec, _ := strconv.ParseInt(inputTime, 10, 64)
		lastTime = time.UnixMilli(msec)
	} else {
		lastTime = time.Now()
	}

	videoService := GetVideo()
	if token, ok = c.Get(jwt.IdentityKey); ok == true {
		userId = token.(*model.User).Id
		feed, nextTime, err = videoService.Feed(userId, lastTime)
	} else {
		feed, nextTime, err = videoService.Feed(-1, lastTime)
	}

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
	var token interface{}
	var userId int64
	var title string
	var data *multipart.FileHeader
	var ok bool = false
	if token, ok = c.Get(jwt.IdentityKey); ok == true {
		userId = token.(*model.User).Id
		title = c.Query("title")
		data, _ = c.FormFile("data")
	} else {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "Publish Get info failed",
		})
		return
	}
	videoService := service.VideoServiceImpl{}
	if err := videoService.Publish(userId, title, data); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "Publish failed",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "Publish suc",
	})
}

// 获取用户上传列表
func GetPublishList(ctx context.Context, c *app.RequestContext) {
	var token interface{}
	var userId int64
	var ok bool = false
	if token, ok = c.Get(jwt.IdentityKey); ok == true {
		userId = token.(*model.User).Id
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Please login first"},
		})
		return
	}
	strOtherUserId := c.Query("user_id")
	otherUserId, _ := strconv.ParseInt(strOtherUserId, 10, 64)
	videoService := GetVideo()
	publishList, err := videoService.GetPublishVideoList(userId, otherUserId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Get PublishList failed"},
		})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response:  Response{StatusCode: 1, StatusMsg: "Get PublishList suc"},
			VideoList: publishList,
		})
	}
}
