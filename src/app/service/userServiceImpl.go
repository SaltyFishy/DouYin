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
func (usi *UserServiceImpl) GetUserList() []model.User {
	users, err := model.GetUserList()
	if err != nil {
		log.Println(err.Error())
		return users
	}
	return users
}

// Service 根据名称获取用户
func (usi *UserServiceImpl) GetUserByUsername(username string) model.User {
	user, err := model.GetUserByName(username)
	if err != nil {
		log.Println(err.Error())
		return user
	}
	return user
}

// Service 根据id获取用户
func (usi *UserServiceImpl) GetUserById(id int64) model.User {
	user, err := model.GetUserById(id)
	if err != nil {
		log.Println(err.Error())
		return user
	}
	return user
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

// 不需要登录情况下,根据id获得User对象
// 开发ing-------------------------------------------------------------------------------------
func (usi *UserServiceImpl) GetUserWithoutId(id int64) (User, error) {
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
	log.Println("Query User Success")
	//followCount, _ := usi.GetFollowingCnt(id)
	//if err != nil {
	//	log.Println("Err:", err.Error())
	//}
	//followerCount, _ := usi.GetFollowerCnt(id)
	//if err != nil {
	//	log.Println("Err:", err.Error())
	//}
	user = User{
		Id:            id,
		Name:          modelUser.Username,
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
	return user, nil
}

// 已登录情况下,根据id获得User对象
// 开发ing-------------------------------------------------------------------------------------
func (usi *UserServiceImpl) GetUserByCurId(curId int64) (User, error) {
	//user, err := model.GetUserById(id)
	//if err != nil {
	//	log.Println(err.Error())
	//	return user, err
	//}
	//return user, nil
	return User{}, nil
}
