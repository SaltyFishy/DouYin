package service

import (
	"DouYin/src/app/model"
)

// controller返回的User结构体
type User struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

// UserService接口
type UserService interface {
	/*
		登录状态
	*/
	// 获得全部User对象
	GetUserList() []model.User

	// 根据username获得User对象
	GetUserByUsername(username string) model.User

	// 根据id获得User对象
	GetUserById(id int64) model.User

	// 将User插入表内
	InsertUser(User *model.User) bool
	/*
		未登录状态
	*/
	// 未登录情况下,根据id获得User对象
	GetUserWithoutId(id int64) (User, error)

	// 已登录情况下,根据curId获得User对象
	GetUserByCurId(curId int64) (User, error)
}
