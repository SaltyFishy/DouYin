package model

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type Comment struct {
	Id         int64
	UserId     int64
	VideoId    int64
	Content    string
	CreateTime time.Time
	Cancel     int8
}

func (Comment) TableName() string {
	return "comments"
}

// 查找是否有过评论记录， 0没有找到，1找到了
func FindCommentByUserIdAndVideoId(userId int64, videoId int64) (int8, error) {
	var comment Comment
	if err := Db.Model(&Comment{}).Where("user_id = ? AND video_id = ?", userId, videoId).First(&comment).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println(err.Error())
		return 0, err
	}
	return 1, nil
}

// 创建评论,1表示评论存在，0表示评论不存在
func CreateComment(userId int64, videoId int64, content string) error {
	comment := Comment{
		UserId:     userId,
		VideoId:    videoId,
		Content:    content,
		CreateTime: time.Now(),
		Cancel:     1,
	}
	if err := Db.Model(&Comment{}).Create(&comment).Error; err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

// 获取某个视频的评论
func GetCommentList(videoId int64) ([]Comment, error) {
	var commentList []Comment
	if err := Db.Model(&Comment{}).Where("video_id = ?", videoId).Find(&commentList).Error; err != nil {
		log.Println("Get commentList Error on commentModel")
		return nil, err
	}
	return commentList, nil
}
