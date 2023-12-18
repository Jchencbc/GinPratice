package routers

import (
	"ginPratice/controllers/account"
	"ginPratice/middlewares"

	"github.com/gin-gonic/gin"
)

func AccountRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/account")
	apiRouters.Use(middlewares.JwtVerify)
	{
		apiRouters.POST("/refresh", account.AccountController{}.TokenRefresh) // token刷新
		apiRouters.GET("/userInfo", account.AccountController{}.UserInfo)     // 用户信息获取
	}
}
