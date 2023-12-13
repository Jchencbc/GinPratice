package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"str": "这是一个测试接口",
			})
		})
		// apiRouters.POST("/submit", submitEndpoint)
		// apiRouters.POST("/read", readEndpoint)
	}
}
