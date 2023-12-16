package account

import (
	"ginPratice/middlewares"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
}

func (con AccountController) CreateAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"str": "这是一个用户创建的接口",
	})
}

func (con AccountController) Login(c *gin.Context) {
	//校验密码逻辑

	c.JSON(http.StatusOK, gin.H{
		"token": middlewares.GenerateToken(&middlewares.UserClaims{
			ID:             "001",
			Name:           "张三",
			StandardClaims: jwt.StandardClaims{},
		}),
	})
}

func (con AccountController) UserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	claims := user.(*middlewares.UserClaims)
	c.JSON(http.StatusOK, gin.H{
		"name": claims.Name,
	})
}
