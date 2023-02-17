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
	// 获得全部User对象
	GetUserList() ([]model.User, error)

	// 根据username获得User对象
	GetUserByUsername(username string) (model.User, error)

	// 根据id获得User对象
	GetUserById(id int64) (model.User, error)

	// 将User插入表内
	InsertUser(User *model.User) bool

	// 封装model.User -> service.User
	GetServiceUserById(id int64) (User, error)
}
