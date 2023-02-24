package model

import (
	"log"
)

type User struct {
	Id       int64
	Username string
	Password string
}

func (User) TableName() string {
	return "users"
}

// 获取所有用户
func GetUserList() ([]User, error) {
	users := []User{}
	if err := Db.Find(&users).Error; err != nil {
		log.Println(err.Error())
		return users, err
	}
	return users, nil
}

// 根据名称获取用户
func GetUserByName(username string) (User, error) {
	user := User{}
	if err := Db.Where("username = ?", username).Find(&user).Error; err != nil {
		log.Println(err.Error())
		return user, err
	}
	return user, nil
}

// 根据Id获取用户
func GetUserById(id int64) (User, error) {
	user := User{}
	if err := Db.Where("id = ?", id).Find(&user).Error; err != nil {
		log.Println(err.Error())
		return user, err
	}
	return user, nil
}

// 插入新用户
func InsertUser(user *User) bool {
	if err := Db.Create(&user).Error; err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

// 检查用户是否合法
func CheckUser(username string, password string) ([]*User, error) {
	user := make([]*User, 0)
	if err := Db.Where("username = ? AND password = ?", username, password).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
