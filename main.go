package main

import (
	"fmt"
	"go-gin-smallproject/pkg/setting"
	"go-gin-smallproject/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort), // 监听的 TCP 地址，格式为:8000
		Handler:        router,                               // http 句柄，实质为ServeHTTP，用于处理程序响应 HTTP 请求
		ReadTimeout:    setting.ReadTimeout,                  // 允许读取的最大时间
		WriteTimeout:   setting.WriteTimeout,                 // 允许读取请求头的最大时间
		MaxHeaderBytes: 1 << 20,                              // 请求头的最大字节数
	}
	s.ListenAndServe()

}
