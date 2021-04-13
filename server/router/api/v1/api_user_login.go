package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-start/server/config"
	"github.com/lazyfury/go-web-start/server/middleware"
	"github.com/lazyfury/go-web-start/server/model"
	"github.com/lazyfury/go-web-template/response"
	"github.com/lazyfury/go-web-template/tools"
)

func login(g *gin.RouterGroup) {

	// login := g.Group("/login")

	// login.POST("", doLogin)

	// login.GET("/init_admin", initAdmin)

}

func doLogin(c *gin.Context) {
	var u = &struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(u); err != nil {
		response.Error("参数错误")
	}

	u.UserName = strings.Trim(u.UserName, " ")
	if u.UserName == "" {
		response.Error("用户昵称不可空")
	}

	u.Password = strings.Trim(u.Password, " ")
	if u.Password == "" {
		response.Error("用户密码不可空")
	}

	user := model.User{Name: u.UserName}

	// err := user.HasUser()
	// if err != nil {
	// 	response.Error("用户不存在")
	// }
	log.Print(config.Global.Sha1)
	password := config.Global.Sha1.EnCode(u.Password)
	if user.Password == password {
		str, _ := middleware.CreateToken(user)
		c.JSON(http.StatusOK, response.JSONSuccess(
			"",
			str,
		))
		return
	}
	response.Error("密码错误")
}

func initAdmin(c *gin.Context) {

	ip := c.ClientIP()
	ip = strings.Split(ip, ":")[0]
	if ip != "127.0.0.1" {
		response.Error(ip)
	}

	secret := c.Query("secret")
	if secret != "fqEeEPlgFECywkwqVMoCEmBzmRmFPZwt" {
		response.Error("密钥错误")
	}
	db := model.DB

	a := &model.User{Name: "admin", IsAdmin: 1}
	if findAdmin := db.Where(a).Find(a).RowsAffected; findAdmin >= 1 {
		a.Password = config.Global.Sha1.AesDecryptCFB(a.Password)
		c.JSON(http.StatusOK, response.JSONSuccess("", a))
		return
	}
	pwd := tools.RandStringBytes(16)

	admin := &model.User{
		Name:     "admin",
		Password: config.Global.Sha1.EnCode(pwd),
		IsAdmin:  1,
	}

	if err := db.Save(admin).Error; err != nil {
		response.Error(err)
	}
	admin.Password = config.Global.Sha1.AesDecryptCFB(admin.Password)
	c.JSON(http.StatusOK, response.JSONSuccess("", admin))
}
