package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/iamzhiyudong/go-gin-example/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int() //将字符串转换为指定类型
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}
