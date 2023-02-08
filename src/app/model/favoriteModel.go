package model

import (
	"DouYin/src/conf"
	"errors"
	"gorm.io/gorm"
	"log"
)

type Favorite struct {
	Id      int64
	UserId  int64
	VideoId int64
	Cancel  int8 // 1为点赞，2为取消点赞
}

func (Favorite) TableName() string {
	return "favorites"
}

// 查找是否有过点赞记录， 0没有找到，1找到了
func FindFavoriteByUserIdAndVideoId(userId int64, videoId int64) (int8, error) {
	var favorite Favorite
	if err := Db.Model(&Favorite{}).Where("user_id = ? AND video_id = ?", userId, videoId).First(&favorite).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println(err.Error())
		return 0, err
	}
	log.Println(favorite)
	return 1, nil
}

// 创建点赞记录
func CreateFavorite(userId int64, videoId int64) error {
	favorite := Favorite{
		UserId:  userId,
		VideoId: videoId,
		Cancel:  1,
	}
	if err := Db.Model(&Favorite{}).Create(&favorite).Error; err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

// 根据userId，videoId,actionType点赞或取消
func UpdateFavorite(userId int64, videoId int64, actionType int32) error {
	if err := Db.Model(&Favorite{}).Where(map[string]interface{}{"user_id": userId, "video_id": videoId}).
		Update("cancel", actionType).Error; err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

// 根据userId查询所属点赞全部videoId
func GetFavoriteVideoIdList(userId int64) ([]int64, error) {
	var FavoriteVideoIdList []int64
	if err := Db.Model(&Favorite{}).Where(map[string]interface{}{"user_id": userId, "cancel": conf.Favorite}).
		Pluck("video_id", &FavoriteVideoIdList).Error; err != nil {
		if "record not found" == err.Error() {
			log.Println("there are no likeVideoId")
			return FavoriteVideoIdList, nil
		} else {
			log.Println(err.Error())
			return FavoriteVideoIdList, err
		}
	}
	return FavoriteVideoIdList, nil
}

// 获取视频被点赞次数
func GetFavoriteCount(videoId int64) (int64, error) {
	var counter int64 = -1
	if err := Db.Model(&Favorite{}).Where("video_id = ? AND cancel = ?", videoId, 1).Find(&counter).Error; err != nil {
		log.Println(err.Error())
		return counter, err
	}
	return counter, nil
}
