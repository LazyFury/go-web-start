package utils

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-template/response"
)

// Error Error
func Error(err interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		response.LogError(fmt.Sprintf("%s %d", file, line))
	}
	panic(err)
}

// Recover 使用defer调用阻止panic中止程序
func Recover(c *gin.Context) {
	if r := recover(); r != nil {
		result := response.ParseError(r)

		var code = c.Writer.Status()
		result.Code = response.ErrCode(code)

		//返回内容
		if ReqFromHTML(c) {
			c.HTML(code, "err/error.html", result)
		} else {
			c.JSON(code, result)
		}
		// c.JSON(code, result)
		response.LogError(fmt.Sprintf("URL:%s ;\nErr: %v", c.Request.URL.RequestURI(), result))

		// "打断response继续写入内容"
		c.Abort()
	}
}
