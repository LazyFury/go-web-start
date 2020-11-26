package api

import (
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/gin-gonic/gin"
)

var modelAd model.Ad
var modelAdGroup model.AdGroup
var modelAdEvent model.AdEvent

func ad(g *gin.RouterGroup) {
	modelAd.BaseControll.Model = &modelAd
	modelAd.Install(g, "/ads")
}

func adGroup(g *gin.RouterGroup) {
	modelAdGroup.BaseControll.Model = &modelAdGroup
	modelAdGroup.Install(g, "/ad-groups")
}

func adEvent(g *gin.RouterGroup) {
	modelAdEvent.BaseControll.Model = &modelAdEvent
	modelAdEvent.Install(g, "/ad-events")
}
