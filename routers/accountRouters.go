package routers

import (
	"fmt"
	"ginPratice/controllers/account"
	"ginPratice/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func initMiddlewareTest(ctx *gin.Context) {
	fmt.Println("中间件测试")
	ctx.Set("username", "张三") //中间件和controller传值

	ctxCopy := ctx.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done!" + ctxCopy.Request.URL.Path)
	}() // 中间件使用goroutine

	ctx.Next() //经行后续请求
	// ctx.Abort()//中止后续请求
	fmt.Println("中间件测试")

}

func AccountRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/account")
	apiRouters.Use(middlewares.InitMiddleware)
	{
		apiRouters.POST("/login", initMiddlewareTest, func(ctx *gin.Context) {

			fmt.Println("login接口使用")
			username, _ := ctx.Get("username") //接受中间件传递的值
			v, ok := username.(string)
			if ok {
				ctx.JSON(http.StatusOK, gin.H{
					"str": v,
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"error": "error",
				})
			}
		})
		apiRouters.POST("/create", account.AccountController{}.CreateAccount)       // 创建用户
		apiRouters.POST("/token/refresh", account.AccountController{}.TokenRefresh) // token刷新
	}
}
