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
		apiRouters.GET("/userInfo", account.AccountController{}.UserInfo) // 用户信息获取
	}
}
