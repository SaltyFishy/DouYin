package service

import (
	"DouYin/src/app/model"
	"errors"
	"gorm.io/gorm"
)

type FavoriteServiceImpl struct {
}

// 点赞行为
func (fsi *FavoriteServiceImpl) FavoriteAction(userId int64, videoId int64, actionType int32) error {
	if actionType == 2 {
		actionType = 0
	}
	if _, err := model.FindFavoriteByUserIdAndVideoId(userId, videoId); errors.Is(err, gorm.ErrRecordNotFound) {
		err := model.CreateFavorite(userId, videoId)
		if err != nil {
			return err
		}
	} else {
		err = model.UpdateFavorite(userId, videoId, actionType)
		if err != nil {
			return err
		}
	}
	return nil
}

// 获取当前用户的所有点赞视频，调用videoService的方法
func (fsi *FavoriteServiceImpl) GetFavoriteList(userId int64) ([]Video, error) {
	videoIdList, _ := model.GetFavoriteVideoIdList(userId)
	var videoList = []Video{}
	usi := UserServiceImpl{}
	for _, id := range videoIdList {
		video, _ := model.GetVideoByVideoId(id)
		Author, _ := usi.GetServiceUserById(video.AuthorId)
		favorite, _ := model.FindFavoriteByUserIdAndVideoId(userId, video.AuthorId)
		favoriteCount, _ := model.GetFavoriteCount(id)
		isFavorite := false
		if favorite == 1 {
			isFavorite = true
		}
		commentCount, _ := model.CountComment(id)
		data := Video{
			Id:            video.Id,
			Author:        Author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			CommentCount:  commentCount,
			FavoriteCount: favoriteCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		}
		videoList = append(videoList, data)
	}
	return videoList, nil
}
