package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(ctx *gin.Context) {
	fmt.Println("局部中间件测试")
	ctx.Next() //经行后续请求
	// ctx.Abort()//中止后续请求
	fmt.Println("局部件测试")
}
