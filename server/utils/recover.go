package utils

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/Treblex/go-web-template/response"
	"github.com/gin-gonic/gin"
)

// Error Error
func Error(err interface{}) {
	log.Println(runtime.Caller(1))
	panic(err)
}

// Recover 使用defer调用阻止panic中止程序
func Recover(c *gin.Context) {
	if r := recover(); r != nil {
		result := response.ParseError(r)
		var code = http.StatusOK
		// if http.StatusText(int(result.Code)) != "" {
		// 	code = int(result.Code)
		// }
		//返回内容
		if ReqFromHTML(c) {
			c.HTML(code, "err/error.html", result)
		} else {
			c.JSON(code, result)
		}
		// c.JSON(code, result)
		response.LogError(fmt.Sprintf("URL:%s ;\nErr: %v", c.Request.URL.RequestURI(), result))

		// "打断response继续写入内容"
		// c.AbortWithStatus(http.StatusInternalServerError)
		c.Abort()

	}
}
