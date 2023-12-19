package main

import (
	"ginPratice/controllers/account"
	"ginPratice/middlewares"
	"ginPratice/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //默认路由引擎

	// db := utils.InitDB()
	// db.AutoMigrate(&models.User{}, &models.Message{}, &models.UserInfo{})

	router.Use(middlewares.InitMiddleware) //全局中间件配置
	router.Use(middlewares.Cors())         //解决跨域
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "你好",
			"data": "lala",
		})
	}) // 测试

	router.POST("/account/create", account.AccountController{}.CreateAccount) // 创建用户
	router.POST("/account/login", account.AccountController{}.Login)          // token获取

	// 用户路由注册
	routers.AccountRoutersInit(router)
	// api应用路由注册
	routers.ApiRoutersInit(router)

	router.Run(":8000")
}
