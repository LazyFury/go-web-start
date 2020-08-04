package api

import (
	"EK-Server/model"

	"github.com/labstack/echo/v4"
)

var modelAd model.Ad
var modelAdGroup model.AdGroup
var modelAdEvent model.AdEvent

func ad(g *echo.Group) {
	modelAd.BaseControll.Model = &modelAd
	modelAd.Install(g, "/ads")
}

func adGroup(g *echo.Group) {
	modelAdGroup.BaseControll.Model = &modelAdGroup
	modelAdGroup.Install(g, "/ad-groups")
}

func adEvent(g *echo.Group) {
	modelAdEvent.BaseControll.Model = &modelAdEvent
	modelAdEvent.Install(g, "/ad-events")
}
