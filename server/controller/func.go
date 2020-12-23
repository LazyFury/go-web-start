package controller

import (
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/Treblex/go-web-template/controller"
	"github.com/Treblex/go-web-template/xmodel"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// auth 还没想好参数需要什么，晚点儿写
func auth(userIDField string, someconfig ...string) controller.Auth {
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

func defaultAuth() controller.Auth {
	return auth("user_id")
}
