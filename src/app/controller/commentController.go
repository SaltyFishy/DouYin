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

type CommentResponse struct {
	Response
	CommentList []service.Comment `json:"comment_list"`
}

type CommentActionResponse struct {
	Response
	Comment service.Comment `json:"comment"`
}

func CommentAction(ctx context.Context, c *app.RequestContext) {
	var token interface{}
	var userId, videoId, actionType int64
	var ok bool = false
	if token, ok = c.Get(jwt.IdentityKey); ok == true {
		userId = token.(*model.User).Id
		strVideoId := c.Query("video_id")
		videoId, _ = strconv.ParseInt(strVideoId, 10, 64)
		strActionType := c.Query("action_type")
		actionType, _ = strconv.ParseInt(strActionType, 10, 64)
	} else {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Token Error"},
		})
		return
	}
	commentService := service.CommentServiceImpl{}
	if actionType == 1 {
		commentText := c.Query("comment_text")
		comment, err := commentService.CreateComment(userId, videoId, commentText)
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 1, StatusMsg: "插入评论失败"},
			})
		} else {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 0, StatusMsg: "插入评论成功"},
				Comment: comment,
			})
		}
	} else if actionType == 2 {
		strCommentId := c.Query("comment_id")
		commentId, _ := strconv.ParseInt(strCommentId, 10, 64)
		err := commentService.DeleteComment(userId, videoId, commentId)
		if err != nil {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 1, StatusMsg: "评论不存在或已删除"},
			})
		} else {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 0, StatusMsg: "删除成功"},
			})
		}
	}
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
			Response:    Response{StatusCode: 0, StatusMsg: "Get CommentList suc"},
			CommentList: commentList,
		})
	}
}
