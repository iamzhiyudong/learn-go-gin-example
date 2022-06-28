package app

import (
	"github.com/astaxie/beego/validation"

	"github.com/iamzhiyudong/go-gin-example/pkg/logging"
)

// 路由参数错误处理
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}
