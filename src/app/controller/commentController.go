package controller

import (
	"DouYin/src/app/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
	"strconv"
)

type CommentResponse struct {
	Response
	CommentList []service.Comment `json:"comment_list"`
}

func CommentAction(ctx context.Context, c *app.RequestContext) {

}

func GetCommentList(ctx context.Context, c *app.RequestContext) {
	strVideoId := c.Query("video_id")
	videoId, _ := strconv.ParseInt(strVideoId, 10, 64)
	commentService := service.CommentServiceImpl{}
	commentList, err := commentService.GetCommentList(videoId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, CommentResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Get CommentList failed"},
		})
	} else {
		c.JSON(http.StatusOK, CommentResponse{
			Response:    Response{StatusCode: 1, StatusMsg: "Get CommentList suc"},
			CommentList: commentList,
		})
	}
}
