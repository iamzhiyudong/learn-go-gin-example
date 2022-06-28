package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iamzhiyudong/go-gin-example/pkg/e"
	"github.com/iamzhiyudong/go-gin-example/pkg/util"
)

// JWT 中间件，处理 JWT 的校验和解析
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")

		if token == "" {
			code = e.ERROR_AUTH
		} else {
			// 解析 token 字符串
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort() // 中断请求
			return
		}

		c.Next() // 放行
	}
}
