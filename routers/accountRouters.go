package routers

import (
	"ginPratice/controllers/account"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AccountRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/account")
	{
		apiRouters.GET("/login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"str": "这是一个登录测试接口",
			})
		})
		apiRouters.GET("/submit", account.AccountController{}.SubmitEndpoint)
		// apiRouters.POST("/read", readEndpoint)
	}
}
