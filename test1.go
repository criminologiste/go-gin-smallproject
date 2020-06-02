package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-smallproject/pkg/setting"
	"net/http"
)

func main() {
	// Default返回一个引擎实例，该引擎实例已经附加了日志记录器和恢复中间件。
	// 返回 Gin 的type Engine struct{...}，里面包含RouterGroup，相当于创建一个路由Handlers，可以后期绑定各类的路由规则和函数、中间件等
	router := gin.Default()
	//router.GET 创建不同的 HTTP 方法绑定到Handlers中，也支持 POST、PUT、DELETE、PATCH、OPTIONS、HEAD 等常用的 Restful 方法
	//Context是gin中的上下文，它允许我们在中间件之间传递变量、管理流、验证 JSON 请求、响应 JSON 请求等，在gin中包含大量Context的方法，
	//例如我们常用的DefaultQuery、Query、DefaultPostForm、PostForm等等
	router.GET("/Test", func(c *gin.Context) {
		// gin.H 就是一个map[string]interface{}
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort), // 监听的 TCP 地址，格式为:8000
		Handler:        router,                               // http 句柄，实质为ServeHTTP，用于处理程序响应 HTTP 请求
		ReadTimeout:    setting.ReadTimeout,                  // 允许读取的最大时间
		WriteTimeout:   setting.WriteTimeout,                 // 允许读取请求头的最大时间
		MaxHeaderBytes: 1 << 20,                              // 请求头的最大字节数
	}
	s.ListenAndServe()

}
