package wechat

import (
	"EK-Server/util"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

var (
	getCurrentSelfmenuInfo string = "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=%s"
)

func getMenu(c echo.Context) (err error) {
	accessToken, err := wechat.GetAccessToken()
	res, err := http.Get(fmt.Sprintf(getCurrentSelfmenuInfo, accessToken))
	if err != nil {
		return util.JSONErr(c, nil, fmt.Sprintf("%s", err))
	}
	m := map[string]interface{}{}
	if err := json.NewDecoder(res.Body).Decode(&m); err != nil {
		return util.JSONErr(c, nil, fmt.Sprintf("%s", err))
	}
	return util.JSONSuccess(c, m, "")
}
