package service

import (
	"DouYin/src/app/model"
	"log"
)

// UserServiceImpl结构
type UserServiceImpl struct {
	FavoriteService
}

// Service 获取所有成员
func (usi *UserServiceImpl) GetUserList() ([]model.User, error) {
	users, err := model.GetUserList()
	if err != nil {
		log.Println(err.Error())
		return []model.User{}, err
	}
	return users, nil
}

// Service 根据名称获取用户
func (usi *UserServiceImpl) GetUserByUsername(username string) (model.User, error) {
	user, err := model.GetUserByName(username)
	if err != nil {
		log.Println(err.Error())
		return model.User{}, err
	}
	return user, nil
}

// Service 根据id获取用户
func (usi *UserServiceImpl) GetUserById(id int64) (model.User, error) {
	user, err := model.GetUserById(id)
	if err != nil {
		log.Println(err.Error())
		return model.User{}, err
	}
	return user, nil
}

// Service 插入新用户
func (usi *UserServiceImpl) InsertUser(user *model.User) bool {
	flag := model.InsertUser(user)
	if flag == false {
		log.Println("insert failed")
		return flag
	}
	return flag
}

// model.User -> service.User
func (usi *UserServiceImpl) GetServiceUserById(id int64) (User, error) {
	user := User{
		Id:            0,
		Name:          "",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
	modelUser, err := model.GetUserById(id)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User Not Found")
		return user, err
	}
	user = User{
		Id:            id,
		Name:          modelUser.Username,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
	return user, nil
}
