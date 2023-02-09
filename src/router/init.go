package router

import (
	"DouYin/src/app/controller"
	"DouYin/src/app/middleware"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
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
		apiRouter.POST("/user/login/", middleware.JwtMiddleware.LoginHandler)
		// 用户信息接口
		apiRouter.GET("/user/", middleware.JwtMiddleware.MiddlewareFunc(), controller.UserInfo)
		// 视频投稿
		apiRouter.POST("/publish/action/", middleware.JwtMiddleware.MiddlewareFunc(), func(ctx context.Context, c *app.RequestContext) {
			c.String(consts.StatusOK, "publish_action")
		})
		// 发布列表
		apiRouter.GET("/publish/list/", middleware.JwtMiddleware.MiddlewareFunc(), controller.GetPublishList)

		/*
			互动接口
		*/
		// 点赞
		apiRouter.POST("/favorite/action/", middleware.JwtMiddleware.MiddlewareFunc(), controller.FavoriteAction)
		// 喜欢列表
		apiRouter.GET("/favorite/list/", middleware.JwtMiddleware.MiddlewareFunc(), controller.GetFavoriteList)
		// 进行评论
		apiRouter.POST("/comment/action/", middleware.JwtMiddleware.MiddlewareFunc(), controller.CommentAction)
		// 视频所有评论
		apiRouter.GET("/comment/list/", middleware.JwtMiddleware.MiddlewareFunc(), controller.GetCommentList)
	}
}
