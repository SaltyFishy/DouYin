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
}

func (Comment) TableName() string {
	return "comments"
}

// 统计评论数
func CountComment(videoId int64) (int64, error) {
	var cnt int64
	if err := Db.Model(&Comment{}).Where("video_id = ?", videoId).Count(&cnt).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println(err.Error())
		return 0, err
	}
	return cnt, nil
}

// 创建评论,1表示评论存在，0表示评论不存在
func CreateComment(userId int64, videoId int64, content string) (Comment, error) {
	comment := Comment{
		UserId:     userId,
		VideoId:    videoId,
		Content:    content,
		CreateTime: time.Now(),
	}
	if err := Db.Model(&Comment{}).Create(&comment).Error; err != nil {
		log.Println(err.Error())
		return Comment{}, err
	}
	return comment, nil
}

// 删除评论
func DeleteComment(userId int64, videoId int64, commentId int64) error {
	var comment Comment
	if err := Db.Model(&Comment{}).Where("user_id = ? AND video_id = ? AND id = ?", userId, videoId, commentId).Find(&comment).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println(err.Error())
		return err
	}
	Db.Model(&Comment{}).Delete(&comment)
	return nil
}

// 获取某个视频的评论
func GetCommentList(videoId int64) ([]Comment, error) {
	var commentList []Comment
	if err := Db.Model(&Comment{}).Where("video_id = ?", videoId).Find(&commentList).Error; err != nil {
		log.Println("Get commentList Error on commentModel")
		return []Comment{}, err
	}
	return commentList, nil
}
