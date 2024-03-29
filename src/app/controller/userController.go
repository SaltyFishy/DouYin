package controller

import (
	"DouYin/src/app/middleware/jwt"
	"DouYin/src/app/model"
	"DouYin/src/app/service"
	"DouYin/src/util"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// 公共返回体
type Response struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
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
	var registerStruct struct {
		Username string `form:"username" json:"username" query:"username" vd:"(len($) > 0 && len($) < 32); msg:'Illegal format'"`
		Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 32); msg:'Illegal format'"`
	}

	if err := c.BindAndValidate(&registerStruct); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "Bad Request",
		})
		return
	}

	usi := service.UserServiceImpl{}

	user, err := usi.GetUserByUsername(registerStruct.Username)

	if errors.Is(err, gorm.ErrRecordNotFound) {

	} else if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "Get User Err",
		})
		return
	}

	if user.Username == registerStruct.Username {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "User already exist",
			},
		})
		return
	} else {
		encodePassword := util.MD5(registerStruct.Password)

		newUser := model.User{
			Username: registerStruct.Username,
			Password: encodePassword,
		}
		if usi.InsertUser(&newUser) != true {
			log.Println("Insert Data Fail")
		}
		u, err := usi.GetUserByUsername(registerStruct.Username)

		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "Get User Error After Insert",
			})
			return
		}

		token := registerStruct.Username + encodePassword

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

	user, err := usi.GetUserByUsername(username)

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "Get User Error",
		})
		return
	}

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
	var token interface{}
	var userId int64
	var ok bool = false

	if token, ok = c.Get(jwt.IdentityKey); ok == true {
		log.Println(token)
		strUserId := c.Query("user_id")
		userId, _ = strconv.ParseInt(strUserId, 10, 64)
		// userId = token.(*model.User).Id
	} else {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Token Error"},
		})
		return
	}

	log.Printf("userController Get : %v\n", userId)

	usi := service.UserServiceImpl{
		FavoriteService: &service.FavoriteServiceImpl{},
	}

	if u, err := usi.GetServiceUserById(userId); err != nil {
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
