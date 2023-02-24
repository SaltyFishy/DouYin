package service

import (
	"time"
)

// controller返回的video结构体
type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author,omitempty"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title,omitempty"`
}

type VideoService interface {
	// 传入时间戳，用户id，返回视频切片数组，以及本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	Feed(userId int64, lastTime time.Time) ([]Video, time.Time, error)

	// 通过authorId来查询对应用户发布的视频，并返回对应的视频切片数组
	GetPublishVideoList(userId int64, authorId int64) ([]Video, error)

	// 上传视频
	Publish(userId int64, title string, playUrl string, coverUrl string) error
}
