package main

import (
	"DouYin/src/app/middleware"
	"DouYin/src/app/model"
	"DouYin/src/router"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	// 载入依赖
	InitDeps()

	// 中间件
	middleware.InitJwt()

	h := server.Default()

	router.Init(h)

	h.Spin()
}

func InitDeps() {
	//加载数据库
	model.Init()

}
