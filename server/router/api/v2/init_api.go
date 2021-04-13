package v2

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-template/response"
)

// Init 初始化
func Init(g *gin.RouterGroup) {
	api := g.Group("/v2")

	api.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.JSONSuccess("welcome!", nil))
	})
}
