package main

import (
	"ginPratice/controllers/account"
	"ginPratice/middlewares"
	"ginPratice/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()                //默认路由引擎
	router.Use(middlewares.InitMiddleware) //全局中间件配置

	router.POST("/account/create", account.AccountController{}.CreateAccount) // 创建用户
	router.POST("/account/login", account.AccountController{}.Login)          // token获取

	// 用户路由注册
	routers.AccountRoutersInit(router)
	// api应用路由注册
	routers.ApiRoutersInit(router)

	router.Run(":8000")
}
