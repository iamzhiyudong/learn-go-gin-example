package main

import (
	"fmt"
	"net/http"

	"github.com/iamzhiyudong/go-gin-example/pkg/setting"
	"github.com/iamzhiyudong/go-gin-example/routers"
)

func main() {
	// === 添加热更新之前启动形式
	router := routers.InitRouter()

	// 实例化一个 http 客户端结构体，绑定处理 handler
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20, // 请求头的最大字节数
	}
	s.ListenAndServe()

	// === 添加热重启
	// endless.DefaultReadTimeOut = setting.ReadTimeout
	// endless.DefaultWriteTimeOut = setting.WriteTimeout
	// endless.DefaultMaxHeaderBytes = 1 << 20
	// endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	// server := endless.NewServer(endPoint, routers.InitRouter())
	// server.BeforeBegin = func(add string) {
	// 	log.Printf("Actual pid is %d", syscall.Getpid())
	// }

	// err := server.ListenAndServe()
	// if err != nil {
	// 	log.Printf("Server err: %v", err)
	// }
}
