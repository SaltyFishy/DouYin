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

	// 评论操作，添加评论
	CreateComment(userId int64, videoId int64, commentText string) error

	// 评论操作，删除评论
	DeleteComment(userId int64, videoId int64, commentId int64) error

	// 根据视频Id获取该视频的评论数
	GetCommentCountByVideoId(videoId int64) (int64, error)
}
