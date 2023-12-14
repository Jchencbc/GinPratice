package main

import (
	"ginPratice/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //默认路由引擎

	// router.Use() //全局中间件配置

	// 用户路由注册
	routers.AccountRoutersInit(router)

	//
	routers.ApiRoutersInit(router)

	router.Run(":8000")
}
