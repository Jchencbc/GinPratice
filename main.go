package main

import (
	"ginPratice/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 用户路由注册
	routers.AccountRoutersInit(router)

	// 简单的路由组: v2
	routers.ApiRoutersInit(router)

	router.Run(":8000")
}
