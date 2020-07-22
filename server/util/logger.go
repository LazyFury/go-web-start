package util

import (
	"errors"
	"io"
	"log"
	"os"
	"time"
)

var (
	// Logger Logger
	Logger *log.Logger = logNew()
)

// LogNew LogNew
func logNew() (logger *log.Logger) {
	//创建文件夹
	dir, err := GetDir("./", "/log/")
	if err != nil {
		err = errors.New("创建文件夹失败")
		return
	}
	filename := dir + "/" + time.Now().Format("2006年01月02日") + ".log"

	logfile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic("取日志文件失败")
	}
	// 定义多个写入器
	writers := []io.Writer{
		logfile,
		os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	// 创建新的log对象
	logger = log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime|log.Lshortfile)
	// os.Stdout = logfile
	os.Stderr = logfile
	logger.SetOutput(logfile)
	return
}
