package controller

import (
	"DouYin/src/app/middleware/jwt"
	"DouYin/src/app/model"
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

	var token interface{}
	var userId, videoId, actionType int64
	var ok bool = false

	if token, ok = c.Get(jwt.IdentityKey); ok == true {
		userId = token.(*model.User).Id
		strVideoId := c.Query("video_id")
		strActionType := c.Query("action_type")
		videoId, _ = strconv.ParseInt(strVideoId, 10, 64)
		actionType, _ = strconv.ParseInt(strActionType, 10, 64)
	} else {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Token Error"},
		})
		return
	}

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
