package service

import "DouYin/src/app/model"

type CommentServiceImpl struct {
}

func (csi *CommentServiceImpl) GetCommentList(videoId int64) ([]Comment, error) {
	modelCommentList, _ := model.GetCommentList(videoId)
	commentList := make([]Comment, 0, len(modelCommentList))
	usi := VideoServiceImpl{}
	for _, comment := range modelCommentList {
		user, _ := usi.GetUserWithoutId(comment.UserId)
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
