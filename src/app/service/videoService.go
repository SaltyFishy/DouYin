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
	Feed(lastTime time.Time) ([]Video, time.Time, error)

	// 传入视频id获得对应的视频对象
	GetVideo(videoId int64) (Video, error)

	// 通过authorId来查询对应用户发布的视频，并返回对应的视频切片数组
	GetPublishVideoList(authorId int64) ([]Video, error)

	// 通过一个作者id，返回该用户发布的视频id切片数组
	GetVideoIdList(authorId int64) ([]int64, error)
}
