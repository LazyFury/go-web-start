package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-start/server/config"
	"github.com/lazyfury/go-web-template/response"
	"github.com/lazyfury/go-web-template/tools/upload"
)

// var uploader = upload.NewEchoUploader()
var aliUploader = upload.NewAliOssUploader(config.Global.AliOss)

// Init upload
func InitUpload(g *gin.RouterGroup) {
	uploadAPI := g.Group("/upload")

	// base
	uploadAPI.POST("", func(c *gin.Context) {
		url, err := aliUploader.Default(c.Request)
		if err != nil {
			response.Error(err)
		}
		c.JSON(http.StatusOK, response.JSONSuccess("", url))
	})

	// only img
	uploadAPI.POST("/upload-img", func(c *gin.Context) {
		url, err := aliUploader.OnlyAcceptsExt(c.Request, upload.AcceptsImgExt, "image")
		if err != nil {
			response.Error(err)
		}
		c.JSON(http.StatusOK, response.JSONSuccess("上传成功", url))
	})

	// custom dir
	uploadAPI.POST("/upload-head-pic", func(c *gin.Context) {
		url, err := aliUploader.Custom(c.Request, upload.AcceptsImgExt, "head_pic")
		if err != nil {
			response.Error(err)
		}
		c.JSON(http.StatusOK, response.JSONSuccess("上传成功", url))
	})
}
