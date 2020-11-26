package api

import (
	"net/http"

	"github.com/Treblex/go-echo-demo/server/config"
	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/gin-gonic/gin"
)

func configRouter(g *gin.RouterGroup) {
	conf := g.Group("/config")

	conf.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.JSONSuccess("", config.Global))
	})
	conf.POST("/save", writeConfig)
}

// 写配置 TODO:todo
func writeConfig(c *gin.Context) {
	c.JSON(http.StatusOK, utils.JSONSuccess("更新设置成功", nil))
}
