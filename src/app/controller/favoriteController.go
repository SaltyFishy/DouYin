package controller

import (
	"DouYin/src/app/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
	"strconv"
)

type FavoriteResponse struct {
	Response
}

type FavoriteListResponse struct {
	Response
	VideoList []service.Video `json:"video_list"`
}

// 点赞
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	// token := c.Query("token")
	strUserId := c.Query("user_id")
	userId, _ := strconv.ParseInt(strUserId, 10, 64)
	strVideoId := c.Query("video_id")
	strActionType := c.Query("action_type")
	videoId, _ := strconv.ParseInt(strVideoId, 10, 64)
	actionType, _ := strconv.ParseInt(strActionType, 10, 64)
	favorite := service.FavoriteServiceImpl{}
	err := favorite.FavoriteAction(userId, videoId, int32(actionType))
	if err == nil {
		log.Println("favoriteAction suc")
		c.JSON(http.StatusOK, FavoriteResponse{
			Response: Response{StatusCode: 0, StatusMsg: "favourite action success"},
		})
	} else {
		log.Println("favoriteAction failed")
		c.JSON(http.StatusOK, FavoriteResponse{
			Response: Response{StatusCode: 1, StatusMsg: "favourite action fail"},
		})
	}
}

// 获取喜欢列表
func GetFavoriteList(ctx context.Context, c *app.RequestContext) {
	strUserId := c.Query("user_id")
	//token := c.Query("token")
	userId, _ := strconv.ParseInt(strUserId, 10, 64)
	fsi := service.FavoriteServiceImpl{}
	favoriteList, err := fsi.GetFavoriteList(userId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, FavoriteListResponse{
			Response: Response{1, "Get favoriteList failed"},
		})
	} else {
		c.JSON(http.StatusOK, FavoriteListResponse{
			Response:  Response{0, "Get favoriteList suc"},
			VideoList: favoriteList,
		})
	}
}
