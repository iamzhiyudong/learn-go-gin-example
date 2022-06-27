package file

import (
	"io/ioutil"
	"mime/multipart" // 主要实现了 MIME 的 multipart 解析，主要适用于 HTTP 和常见浏览器生成的 multipart 主体
	"os"
	"path"
)

func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

func GetExt(fileName string) string {
	return path.Ext(fileName)
}

func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}
