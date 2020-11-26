package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 自己尝试的cors配置实现
func cosr(c *gin.Context) {
	req := c.Request
	origin := c.Request.Header.Get("Origin")
	if len(origin) == 0 {
		// request is not a CORS request
		return
	}
	host := c.Request.Host

	if origin == "http://"+host || origin == "https://"+host {
		// request is not a CORS request but have origin header.
		// for example, use fetch api
		return
	}

	allowOrigins := []string{"*"}
	inAllow := 0
	for i, o := range allowOrigins {
		if origin == o || o == "*" {
			inAllow = i + 1
		}
	}
	if inAllow == 0 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE,OPTION")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "authorization,token,content-type")

	if req.Method == http.MethodOptions {
		c.Status(http.StatusNoContent)
	}
}
