package service

import "time"

type Comment struct {
	Id         int64 `json:"id"`
	User       User
	Content    string    `json:"content"`
	CreateDate time.Time `json:"create_date"`
}

type CommentService interface {

	// 根据videoId获取所有评论
	GetCommentList(videoId int64) ([]Comment, error)
}
