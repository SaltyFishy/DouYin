package router

import (
	"DouYin/src/app/controller"
	"DouYin/src/app/middleware/jwt"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init(h *server.Hertz) {
	// 路由组
	apiRouter := h.Group("/douyin")
	{
		/*
			基础接口
		*/
		// 视频流接口
		apiRouter.GET("/feed/", controller.Feed)
		// 用户注册接口
		apiRouter.POST("/user/register/", controller.Register)
		// 用户登录接口
		apiRouter.POST("/user/login/", jwt.JwtMiddleware.LoginHandler)
		// 用户信息接口
		apiRouter.GET("/user/", jwt.JwtMiddleware.MiddlewareFunc(), controller.UserInfo)
		// 视频投稿
		apiRouter.POST("/publish/action/", jwt.JwtMiddleware.MiddlewareFunc(), controller.PublishAction)
		// 发布列表
		apiRouter.GET("/publish/list/", jwt.JwtMiddleware.MiddlewareFunc(), controller.GetPublishList)

		/*
			互动接口
		*/
		// 点赞
		apiRouter.POST("/favorite/action/", jwt.JwtMiddleware.MiddlewareFunc(), controller.FavoriteAction)
		// 喜欢列表
		apiRouter.GET("/favorite/list/", jwt.JwtMiddleware.MiddlewareFunc(), controller.GetFavoriteList)
		// 进行评论
		apiRouter.POST("/comment/action/", jwt.JwtMiddleware.MiddlewareFunc(), controller.CommentAction)
		// 视频所有评论
		apiRouter.GET("/comment/list/", jwt.JwtMiddleware.MiddlewareFunc(), controller.GetCommentList)
	}
}
