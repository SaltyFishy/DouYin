package service

import (
	"DouYin/src/app/model"
	"DouYin/src/conf"
	"fmt"
	"log"
	"mime/multipart"
	"sync"
	"time"
)

var wg sync.WaitGroup

type VideoServiceImpl struct {
	UserService
	FavoriteService
	CommentService
}

// Service 获取视频
func (vsi *VideoServiceImpl) Feed(userId int64, lastTime time.Time) ([]Video, time.Time, error) {
	videos := make([]Video, 0, conf.VideoMaxCount)
	fmt.Println(lastTime)
	modelVideos, err := model.GetVideosByLastTime(lastTime)
	if err != nil {
		log.Printf(err.Error())
		return nil, time.Time{}, err
	}
	log.Printf("get lastTime suc")
	fmt.Println(modelVideos)
	err = vsi.copyVideos(userId, &videos, &modelVideos)
	if err != nil {
		log.Printf(err.Error())
		return nil, time.Time{}, err
	}
	log.Printf("copy suc")
	fmt.Println(videos)
	return videos, modelVideos[len(modelVideos)-1].PublishTime, nil
}

// 通过userId来查询对应用户发布的视频，并返回对应的视频切片数组
func (vsi *VideoServiceImpl) GetPublishVideoList(userId int64, authorId int64) ([]Video, error) {
	var modelVideoList []model.Video
	modelVideoList, err := model.GetVideosByAuthorId(authorId)
	if err != nil {
		log.Println(err.Error())
		return []Video{}, err
	}
	videoList := make([]Video, 0, len(modelVideoList))
	err = vsi.copyVideos(userId, &videoList, &modelVideoList)
	if err != nil {
		log.Printf(err.Error())
		return []Video{}, err
	}
	return videoList, nil
}

// 原地拷贝，modelVideos -> videos
func (vsi *VideoServiceImpl) copyVideos(userId int64, videos *[]Video, modelVideos *[]model.Video) error {
	for i, data := range *modelVideos {
		video := Video{Id: int64(i)}
		vsi.generateVideo(userId, &video, &data)
		*videos = append(*videos, video)
	}
	return nil
}

// 原地构造，model.Video -> service.Video
func (vsi *VideoServiceImpl) generateVideo(userId int64, video *Video, data *model.Video) {
	video.Id, video.PlayUrl, video.CoverUrl, video.Title = data.Id, data.PlayUrl, data.CoverUrl, data.Title
	author, err := vsi.GetServiceUserById(data.AuthorId)
	if err != nil {
		video.Author = User{}
		log.Printf("方法vsi.GetServiceUserById 失败：%v", err)
	} else {
		video.Author = author
		log.Printf("方法vsi.GetServiceUserById 成功")
	}

	video.FavoriteCount, err = model.GetFavoriteCount(data.Id)
	if err != nil {
		log.Printf("方法model.GetFavoriteCount 失败：%v", err)
	} else {
		log.Printf("方法model.GetFavoriteCount 成功")
	}

	video.CommentCount, err = vsi.GetCommentCountByVideoId(data.Id)
	if err != nil {
		log.Printf("方法GetCommentCountByVideoId 失败：%v", err)
	} else {
		log.Printf("方法GetCommentCountByVideoId 成功")
	}

	favorite, err := model.FindFavoriteByUserIdAndVideoId(userId, video.Id)
	if err != nil {
		log.Printf("方法FindFavoriteByUserIdAndVideoId 失败：%v", err)
	} else {
		log.Printf("方法FindFavoriteByUserIdAndVideoId 成功")
		if favorite == 1 {
			video.IsFavorite = true
		}
	}
}

func (vsi *VideoServiceImpl) Publish(userId int64, title string, data *multipart.FileHeader) error {

	return nil
}
