package tg

import (
	"EK-Server/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

var (
	rootURL string = "https://api.telegram.org/bot"
	tg      *app   = &app{
		key: "1097553478:AAFdTSRGrSPrXaTRu3eeSzqBFusWPQzZZZo",
	}
)

type (
	app struct {
		key string
	}
)

func (a *app) getMe(c echo.Context) (err error) {
	res, err := http.Get(fmt.Sprintf("%s%s%s", rootURL, a.key, "/getMe"))
	if err != nil {
		log.Println(err)
		return util.JSONErr(c, nil, "请求失败")
	}
	defer res.Body.Close()
	data := map[string]interface{}{}
	_ = json.NewDecoder(res.Body).Decode(&data)

	return util.JSONSuccess(c, data, "获取成功")
}

// Init 初始化
func Init(g *echo.Group) {
	tgRoute := g.Group("tg")
	tgRoute.GET("/getMe", tg.getMe)
}
