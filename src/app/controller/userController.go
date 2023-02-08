package controller

import (
	"DouYin/src/app/model"
	"DouYin/src/app/service"
	"DouYin/src/util"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
	"strconv"
)

// 公共返回体
type Response struct {
	StatusCode int32  "json:`status_code`"          // 状态码，0-成功，其他值-失败
	StatusMsg  string "json:`status_msg,omitempty`" // 返回状态描述
}

// 用户请求返回
type UserResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

// 用户信息返回
type UserInfoResponse struct {
	Response
	User service.User `json:"user"`
}

// 用户注册
func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	usi := service.UserServiceImpl{}

	user := usi.GetUserByUsername(username)

	if user.Username == username {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else if len(username) > 32 {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "length of username should be less than 32"},
		})
	} else if len(password) > 32 {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "length of password should be less than 32"},
		})
	} else {
		newUser := model.User{
			Username: username,
			Password: util.MD5(password),
		}
		if usi.InsertUser(&newUser) != true {
			log.Println("Insert Data Fail")
		}
		u := usi.GetUserByUsername(username)

		encodePassword := util.MD5(password)

		token := username + encodePassword

		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			UserId:   u.Id,
			Token:    token,
		})
	}

}

// 用户登录
func Login(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	encodePassword := util.MD5(password)

	usi := service.UserServiceImpl{}

	user := usi.GetUserByUsername(username)

	if encodePassword == user.Password {
		token := username + encodePassword
		log.Println(token)
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Username or Password Error"},
		})
	}
}

// 用户信息
func UserInfo(ctx context.Context, c *app.RequestContext) {
	userId := c.Query("user_id")
	id, _ := strconv.ParseInt(userId, 10, 64)

	usi := service.UserServiceImpl{
		//	FollowService: &service.FollowServiceImp{},
		FavoriteService: &service.FavoriteServiceImpl{},
	}

	if u, err := usi.GetUserWithoutId(id); err != nil {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User Doesn't Exist"},
		})
	} else {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 0},
			User:     u,
		})
	}
}
