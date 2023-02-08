package router

import (
	"DouYin/src/app/controller"
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
		apiRouter.POST("/user/login/", controller.Login)
		// 用户信息接口
		apiRouter.GET("/user/", controller.UserInfo)
		// 视频投稿
		apiRouter.POST("/publish/action/", func(ctx context.Context, c *app.RequestContext) {
			c.String(consts.StatusOK, "publish_action")
		})
		// 发布列表
		apiRouter.GET("/publish/list/", controller.GetPublishList)

		/*
			互动接口
		*/
		// 点赞
		apiRouter.POST("/favorite/action/", controller.FavoriteAction)
		// 喜欢列表
		apiRouter.GET("/favorite/list/", controller.GetFavoriteList)
		// 进行评论
		apiRouter.POST("/comment/action/", func(ctx context.Context, c *app.RequestContext) {
			c.String(consts.StatusOK, "/comment/action/")
		})
		// 视频所有评论
		apiRouter.GET("/comment/list/", func(ctx context.Context, c *app.RequestContext) {
			c.String(consts.StatusOK, "/comment/list/")
		})
	}
}
