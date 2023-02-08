package service

type FavoriteService interface {
	/*
	   基本操作
	*/
	// 1点赞，2取消点赞。
	FavoriteAction(userId int64, videoId int64, actionType int32) error
	// 获取当前用户的所有点赞视频，调用videoService的方法
	GetFavoriteList(userId int64) ([]Video, error)
}
