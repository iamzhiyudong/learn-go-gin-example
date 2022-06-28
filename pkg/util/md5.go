package util

import (
	"crypto/md5"
	"encoding/hex"
)

// md5 加密函数 - 处理文件名
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
