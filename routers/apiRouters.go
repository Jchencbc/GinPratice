package routers

import (
	"ginPratice/controllers/api"
	"ginPratice/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	apiRouters.Use(middlewares.JwtVerify)
	{
		apiRouters.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"str": "这是一个测试接口",
			})
		})
		apiRouters.POST("/file/upload", api.FileController{}.FileUpload)    //文件上传接口
		apiRouters.GET("/file/download", api.FileController{}.FileDownload) //文件下载接口
		apiRouters.POST("/message/up", api.MessageController{}.MsgUp)       //消息发布
		// apiRouters.POST("/read", readEndpoint)
	}
}
