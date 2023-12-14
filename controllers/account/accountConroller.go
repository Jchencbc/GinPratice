package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
}

func (con AccountController) CreateAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"str": "这是一个用户创建的接口",
	})
}

func (con AccountController) TokenRefresh(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"str": "token 刷新接口",
	})
}
