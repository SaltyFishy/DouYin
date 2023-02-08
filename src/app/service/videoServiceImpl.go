package service

import (
	"DouYin/src/app/model"
	"DouYin/src/conf"
	"fmt"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

type VideoServiceImpl struct {
	UserService
}

// Service 获取视频
func (vsi *VideoServiceImpl) Feed(lastTime time.Time) ([]Video, time.Time, error) {
	videos := make([]Video, 0, conf.VideoMaxCount)
	fmt.Println(lastTime)
	modelVideos, err := model.GetVideosByLastTime(lastTime)
	if err != nil {
		log.Printf(err.Error())
		return nil, time.Time{}, err
	}
	log.Printf("get lastTime suc")
	fmt.Println(modelVideos)
	err = vsi.copyVideos(&videos, &modelVideos)
	if err != nil {
		log.Printf(err.Error())
		return nil, time.Time{}, err
	}
	log.Printf("copy suc")
	fmt.Println(videos)
	return videos, modelVideos[len(modelVideos)-1].PublishTime, nil
}

// Service 根据videoId获取视频对象
func (vsi *VideoServiceImpl) GetVideo(videoId int64) (Video, error) {
	return Video{}, nil
}

// 通过userId来查询对应用户发布的视频，并返回对应的视频切片数组
func (vsi *VideoServiceImpl) GetPublishVideoList(authorId int64) ([]Video, error) {
	var modelVideoList []model.Video
	modelVideoList, err := model.GetVideosByAuthorId(authorId)
	if err != nil {
		log.Println(err.Error())
		return []Video{}, err
	}
	videoList := make([]Video, 0, len(modelVideoList))
	err = vsi.copyVideos(&videoList, &modelVideoList)
	if err != nil {
		log.Printf(err.Error())
		return []Video{}, err
	}
	return videoList, nil
}

// 通过一个作者id，返回该用户发布的视频id切片数组
func (vsi *VideoServiceImpl) GetVideoIdList(authorId int64) ([]int64, error) {
	return []int64{}, nil
}

// 原地拷贝，modelVideos -> videos
func (vsi *VideoServiceImpl) copyVideos(videos *[]Video, modelVideos *[]model.Video) error {
	for i, data := range *modelVideos {
		video := Video{Id: int64(i)}
		vsi.generateVideo(&video, &data)
		*videos = append(*videos, video)
	}
	return nil
}

// 原地构造，model.Video -> service.Video
func (vsi *VideoServiceImpl) generateVideo(video *Video, data *model.Video) {
	//wg.Add(4)
	//var err error
	video.Id, video.PlayUrl, video.CoverUrl, video.Title = data.Id, data.PlayUrl, data.CoverUrl, data.Title
	////插入Author，这里需要将视频的发布者和当前登录的用户传入，才能正确获得isFollow，
	////如果出现错误，不能直接返回失败，将默认值返回，保证稳定
	//go func() {
	//	video.Author, err = videoService.GetUserByIdWithCurId(data.AuthorId, userId)
	//	if err != nil {
	//		log.Printf("方法videoService.GetUserByIdWithCurId(data.AuthorId, userId) 失败：%v", err)
	//	} else {
	//		log.Printf("方法videoService.GetUserByIdWithCurId(data.AuthorId, userId) 成功")
	//	}
	//	wg.Done()
	//}()
	//
	////插入点赞数量，同上所示，不将nil直接向上返回，数据没有就算了，给一个默认就行了
	//go func() {
	//	video.FavoriteCount, err = videoService.FavouriteCount(data.Id)
	//	if err != nil {
	//		log.Printf("方法videoService.FavouriteCount(data.ID) 失败：%v", err)
	//	} else {
	//		log.Printf("方法videoService.FavouriteCount(data.ID) 成功")
	//	}
	//	wg.Done()
	//}()
	//
	////获取该视屏的评论数字
	//go func() {
	//	video.CommentCount, err = videoService.CountFromVideoId(data.Id)
	//	if err != nil {
	//		log.Printf("方法videoService.CountFromVideoId(data.ID) 失败：%v", err)
	//	} else {
	//		log.Printf("方法videoService.CountFromVideoId(data.ID) 成功")
	//	}
	//	wg.Done()
	//}()
	//
	////获取当前用户是否点赞了该视频
	//go func() {
	//	video.IsFavorite, err = videoService.IsFavourite(video.Id, userId)
	//	if err != nil {
	//		log.Printf("方法videoService.IsFavourit(video.Id, userId) 失败：%v", err)
	//	} else {
	//		log.Printf("方法videoService.IsFavourit(video.Id, userId) 成功")
	//	}
	//	wg.Done()
	//}()
	//
	//wg.Wait()
}
