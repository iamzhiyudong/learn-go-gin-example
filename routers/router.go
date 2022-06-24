package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/iamzhiyudong/go-gin-example/pkg/setting"
)

func InitRouter() *gin.Engine {
	// 创建路由 Handlers
	// 后期可以绑定各类的路由规则和函数、中间件等
	// router := gin.Default()

	r := gin.New() // 源码中会调用打印启动日志的函数

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	// 创建不同的 HTTP 方法绑定到Handlers中
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return r
}
