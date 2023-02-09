package controller

import (
	"DouYin/src/app/middleware"
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

	user := usi.GetUserByUsername(registerStruct.Username)

	if user.Username == registerStruct.Username {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "User already exist",
			},
		})
	} else {
		encodePassword := util.MD5(registerStruct.Password)

		newUser := model.User{
			Username: registerStruct.Username,
			Password: encodePassword,
		}
		if usi.InsertUser(&newUser) != true {
			log.Println("Insert Data Fail")
		}
		u := usi.GetUserByUsername(registerStruct.Username)

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
	var token interface{}
	var userId int64
	var ok bool = false

	if token, ok = c.Get(middleware.IdentityKey); ok == true {
		log.Println(token)
		strUserId := c.Query("user_id")
		userId, _ = strconv.ParseInt(strUserId, 10, 64)
		// userId = token.(*model.User).Id
	} else {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Token Error"},
		})

	}

	log.Printf("userController Get : %v\n", userId)

	usi := service.UserServiceImpl{
		//	FollowService: &service.FollowServiceImp{},
		FavoriteService: &service.FavoriteServiceImpl{},
	}

	if u, err := usi.GetUserWithoutId(userId); err != nil {
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
