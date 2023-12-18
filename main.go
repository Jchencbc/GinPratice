package main

import (
	"fmt"
	"ginPratice/controllers/account"
	"ginPratice/middlewares"
	"ginPratice/routers"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

func main() {
	router := gin.Default()                //默认路由引擎
	router.Use(middlewares.InitMiddleware) //全局中间件配置
	router.Use(middlewares.Cors())         //解决跨域
	router.GET("/test", func(ctx *gin.Context) {
		cfg, _ := ini.Load("./config/app.ini")
		host := cfg.Section("Mysql").Key("host").String()
		port := cfg.Section("Mysql").Key("port").String()
		database := cfg.Section("Mysql").Key("database").String()
		username := cfg.Section("Mysql").Key("username").String()
		password := cfg.Section("Mysql").Key("password").String()
		charset := cfg.Section("Mysql").Key("charset").String()
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%v&parseTime=True&loc=Local",
			username,
			password,
			host,
			port,
			database,
			charset)
		fmt.Print(dsn)
		// fmt.Println(dsn)
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
