package main

import (
	"fmt"
	"net/http"

	"github.com/iamzhiyudong/go-gin-example/models"
	"github.com/iamzhiyudong/go-gin-example/pkg/logging"
	"github.com/iamzhiyudong/go-gin-example/pkg/setting"
	"github.com/iamzhiyudong/go-gin-example/routers"
)

func main() {
	// 初始化
	setting.Setup()
	models.Setup()
	logging.Setup()
	// === 添加热更新之前启动形式
	router := routers.InitRouter()

	// 实例化一个 http 客户端结构体，绑定处理 handler
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20, // 请求头的最大字节数
	}
	s.ListenAndServe()

	// === 添加热重启
	// endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	// endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	// endless.DefaultMaxHeaderBytes = 1 << 20
	// endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	// server := endless.NewServer(endPoint, routers.InitRouter())
	// server.BeforeBegin = func(add string) {
	// 	log.Printf("Actual pid is %d", syscall.Getpid())
	// }

	// err := server.ListenAndServe()
	// if err != nil {
	// 	log.Printf("Server err: %v", err)
	// }
}
