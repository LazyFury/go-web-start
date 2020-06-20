package util

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo"
)

var (
	// Logger Logger
	Logger = &LoggerStruct{
		Logger: logNew(),
	}
)

type (
	// LoggerStruct LoggerStruct
	LoggerStruct struct {
		*log.Logger
	}
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
	// log.SetOutput(logfile)
	return
}

// ClientIP 尽最大努力实现获取客户端 IP
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(c echo.Context) string {
	r := c.Request()
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	return strings.TrimSpace(r.RemoteAddr)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandStringBytes 随机字符串
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

//GetDir 创建文件夹
func GetDir(path string, foderName string) (dir string, err error) {
	folder := filepath.Join(path, foderName)
	if _, err = os.Stat(folder); os.IsNotExist(err) {
		err = os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			return
		}
	}
	dir = folder
	return
}
