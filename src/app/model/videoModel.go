package model

import (
	"DouYin/src/conf"
	"log"
	"time"
)

// Video
type Video struct {
	Id          int64     `json:"id,omitempty"`
	AuthorId    int64     `json:"author_id,omitempty"`
	PlayUrl     string    `json:"play_url,omitempty"`
	CoverUrl    string    `json:"cover_url,omitempty"`
	PublishTime time.Time `json:"publish_time,omitempty"`
	Title       string    `json:"title,omitempty"`
}

func (Video) TableName() string {
	return "videos"
}

// 根据作者Id获取Video
func GetVideosByAuthorId(authorId int64) ([]Video, error) {
	video := []Video{}
	if err := Db.Where(&Video{AuthorId: authorId}).Find(&video).Error; err != nil {
		log.Println(err.Error())
		return video, err
	}
	return video, nil
}

// 根据VideoId来获得视频信息
func GetVideoByVideoId(id int64) (Video, error) {
	video := Video{}
	if err := Db.Where("id = ?", id).First(&video).Error; err != nil {
		log.Println(err.Error())
		return video, err
	}
	return video, nil
}

// 获取一个时间之前的一些视频
func GetVideosByLastTime(lastTime time.Time) ([]Video, error) {
	videos := make([]Video, conf.VideoMaxCount)
	if err := Db.Where("publish_time < ?", lastTime).Order("publish_time desc").Limit(conf.VideoMaxCount).Find(&videos).Error; err != nil {
		return videos, err
	}
	return videos, nil
}
