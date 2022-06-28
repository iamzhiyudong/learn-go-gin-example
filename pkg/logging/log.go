package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota // 0
	INFO
	WARNING
	ERROR
	FATAL
)

// 初始化日志保存路径、日志名、日志文件
func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = openLogFile(fileName, filePath)
	if err != nil {
		log.Fatalln(err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// 封装日志打印基本方法
// 调试
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
}

// 信息类
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}

// 警告类
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}

// 错误
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}

// 致命
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v...)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth) // 获取当前堆栈信息
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
