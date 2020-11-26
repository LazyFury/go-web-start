package v2

import (
	"net/http"

	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/gin-gonic/gin"
)

// Init 初始化
func Init(g *gin.RouterGroup) {
	api := g.Group("/v2")

	api.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.JSONSuccess("welcome!", nil))
	})
}
