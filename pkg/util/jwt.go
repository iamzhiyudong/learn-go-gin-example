package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/iamzhiyudong/go-gin-example/pkg/setting"
)

// byte 类型的 jwtSecret
var jwtSecret = []byte(setting.AppSetting.JwtSecret)

// jwt 的声明结构体
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 生成 token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	// 初始化 jwt 的声明结构体
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	// 通过 jwt 结构体生成 token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// token 解析
func ParseToken(token string) (*Claims, error) {
	// 通过 token 字符串、jwtSecret 获取 token 结构体
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) { // 方法内部主要是具体的解码和校验的过程，最终返回*Token
		return jwtSecret, nil
	})

	// 处理解析 token 后的 jwt 实例
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
