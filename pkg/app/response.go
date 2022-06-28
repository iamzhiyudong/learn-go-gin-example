package app

import (
	"github.com/gin-gonic/gin"

	"github.com/iamzhiyudong/go-gin-example/pkg/e"
)

// 路由实例结构体
type Gin struct {
	C *gin.Context
}

// 封装响应处理函数
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})

	return
}
