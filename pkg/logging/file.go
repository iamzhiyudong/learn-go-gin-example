package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/iamzhiyudong/go-gin-example/pkg/file"
	"github.com/iamzhiyudong/go-gin-example/pkg/setting"
)

// 获取日志文件路径
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

// 生成日志文件名
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}

// 打开日志文件
func openLogFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd() // 当前根目录
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := file.CheckPermission(src) // 检查权限
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src) // 是否存在，不存在就创建
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // 打开
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}
