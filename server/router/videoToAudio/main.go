package videoToAudio

import "github.com/labstack/echo"

// Init 初始化
func Init(g *echo.Group) {
	baseURL := "/videoToAudio"
	_ = g.Group(baseURL)
}
 