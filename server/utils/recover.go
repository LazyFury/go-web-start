package utils

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func logError(err interface{}) {
	log.Printf("\n\n\x1b[31m[Custom Debug Result]: %v \x1b[0m\n\n", err)
}

// DuplicateEntryKey DuplicateEntryKey
var DuplicateEntryKey, _ = regexp.Compile(
	`Duplicate entry '(.*)' for key '(.*)\.(.*)'`,
)

// Recover 使用defer调用阻止panic中止程序
func Recover(c *gin.Context) {
	if r := recover(); r != nil {
		result := JSON(http.StatusInternalServerError, "", nil)

		//普通错误
		if err, ok := r.(error); ok {
			logError(err)
			if err == gorm.ErrRecordNotFound {
				result.Code = NotFound
				result.Message = StatusText(result.Code)
			} else {
				result.Message = err.Error()
			}

			if arr := DuplicateEntryKey.FindSubmatch([]byte(err.Error())); len(arr) > 0 {
				result.Message = string(arr[3]) + " 是不可重复字段,已存在相同的数据"
			}
		}
		//错误提示
		if err, ok := r.(string); ok {
			result.Message = err
		}
		//错误码
		if err, ok := r.(ErrCode); ok {
			result.Code = err
			result.Message = StatusText(err)
		} else if err, ok := r.(int); ok {
			result.Message = StatusText(ErrCode(err))
		}
		//完整错误类型
		if data, ok := r.(Result); ok {
			result = data
		}

		var code = http.StatusOK
		// if http.StatusText(int(result.Code)) != "" {
		// 	code = int(result.Code)
		// }
		//返回内容
		// if ReqFromHTML(c) {
		// 	c.HTML(code, "error.html", nil)
		// } else {
		// 	c.JSON(code, result)
		// }
		c.JSON(code, result)
		logError(fmt.Sprintf("URL:%s ;\nErr: %v", c.Request.URL.RequestURI(), result))

		// "打断response继续写入内容"
		c.AbortWithStatus(http.StatusInternalServerError)

	}
}
