package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-start/server/model"
	"github.com/lazyfury/go-web-template/controller"
	xmodel "github.com/lazyfury/go-web-template/model"
	"gorm.io/gorm"
)

// auth 还没想好参数需要什么，晚点儿写
func authWithFilter(userIDField string, someconfig ...string) controller.Auth {
	return func(c *gin.Context, must bool) xmodel.Middleware {
		return func(db *gorm.DB) *gorm.DB {
			var user *model.User
			if must {
				user = model.GetUserOrLogin(c)
			} else {
				user = model.GetUserOrEmpty(c)
			}
			db = db.Where(map[string]interface{}{
				userIDField: user.ID,
			})
			return db
		}
	}
}

func justAuth() controller.Auth {
	return func(c *gin.Context, must bool) xmodel.Middleware {
		return func(db *gorm.DB) *gorm.DB {
			//公开的接口，列表和详情不需要验证
			if c.Request.Method != http.MethodGet {
				model.GetUserOrLogin(c)
			}
			return db
		}
	}
}

func defaultAuth() controller.Auth {
	return justAuth()
}
