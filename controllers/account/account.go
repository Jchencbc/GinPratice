package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
}

func (con AccountController) SubmitEndpoint(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"str": "这是一个用户测试接口",
	})
}
