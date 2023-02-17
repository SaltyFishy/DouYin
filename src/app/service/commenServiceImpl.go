package service

import (
	"DouYin/src/app/model"
	"log"
)

type CommentServiceImpl struct {
}

func (csi *CommentServiceImpl) CreateComment(userId int64, videoId int64, commentText string) error {
	if err := model.CreateComment(userId, videoId, commentText); err != nil {
		return err
	}
	return nil
}

func (csi *CommentServiceImpl) DeleteComment(userId int64, videoId int64, commentId int64) error {
	if err := model.DeleteComment(userId, videoId, commentId); err != nil {
		return err
	}
	return nil
}

func (csi *CommentServiceImpl) GetCommentList(videoId int64) ([]Comment, error) {
	modelCommentList, _ := model.GetCommentList(videoId)
	commentList := make([]Comment, 0, len(modelCommentList))
	vsi := VideoServiceImpl{}
	for _, comment := range modelCommentList {
		user, _ := vsi.GetServiceUserById(comment.UserId)
		data := Comment{
			Id:         comment.Id,
			User:       user,
			Content:    comment.Content,
			CreateDate: comment.CreateTime,
		}
		commentList = append(commentList, data)
	}
	return commentList, nil
}

func (csi *CommentServiceImpl) GetCommentCountByVideoId(videoId int64) (int64, error) {
	var cnt int64
	var err error
	if cnt, err = model.CountComment(videoId); err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return cnt, nil
}
