package jwt

import (
	"DouYin/src/app/model"
	"DouYin/src/conf"
	"DouYin/src/util"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
	"net/http"
	"time"
)

// 公共返回体
type Response struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
}

// 用户请求返回
type UserResponse struct {
	Response
	UserId int64     `json:"user_id,omitempty"`
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "identity"
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "hertz my-jwt",
		Key:           conf.JwtSecret,
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			expire.Format(time.RFC3339)
			//id := GetUserIdFromJwtToken(ctx, c)
			//if id == -1 {
			//	return
			//}
			user, _ := c.Get(JwtMiddleware.IdentityKey)
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 0, StatusMsg: "success"},
				UserId:   user.(*model.User).Id,
				Token:    token,
				Expire:   expire,
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct struct {
				Username string `form:"username" json:"username" query:"username" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
				Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
			}
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			users, err := model.CheckUser(loginStruct.Username, util.MD5(loginStruct.Password))
			if err != nil {
				return nil, err
			}
			if len(users) == 0 {
				return nil, errors.New("user already exists or wrong password")
			}
			c.Set("identity", users[0])
			c.Next(ctx)
			return users[0], nil
		},
		IdentityKey: IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims, _ := JwtMiddleware.GetClaimsFromJWT(ctx, c)
			userMapInterface, _ := claims[IdentityKey].(map[string]interface{})
			return &model.User{
				Id:       int64(userMapInterface["Id"].(float64)),
				Username: userMapInterface["Username"].(string),
				Password: userMapInterface["Password"].(string),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if u, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					IdentityKey: u,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, Response{
				StatusCode: int32(code),
				StatusMsg:  message,
			})
		},
	})
	if err != nil {
		panic(err)
	}

}
