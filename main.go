package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iamzhiyudong/go-gin-example/pkg/setting"
	"github.com/iamzhiyudong/go-gin-example/routers"
)

func main() {
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
	log.Fatal(2, "ListenAndServe: %v", "123")
}
