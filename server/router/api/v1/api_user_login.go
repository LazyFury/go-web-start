package api

import (
	"net/http"
	"strings"

	"github.com/Treblex/go-echo-demo/server/middleware"
	"github.com/Treblex/go-echo-demo/server/model"
	"github.com/Treblex/go-echo-demo/server/utils"
	"github.com/Treblex/go-web-template/tools"
	"github.com/Treblex/simple-daily/utils/sha"
	"github.com/gin-gonic/gin"
)

func login(g *gin.RouterGroup) {

	login := g.Group("/login")

	login.POST("", doLogin)

	login.POST("/reg", modelUser.RegController)

	login.GET("/init_admin", initAdmin)

}

func doLogin(c *gin.Context) {
	var u = &struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(u); err != nil {
		panic("参数错误")
	}

	u.UserName = strings.Trim(u.UserName, " ")
	if u.UserName == "" {
		panic("用户昵称不可空")
	}

	u.Password = strings.Trim(u.Password, " ")
	if u.Password == "" {
		panic("用户密码不可空")
	}

	user := model.User{Name: u.UserName}

	err := user.HasUser()
	if err != nil {
		panic("用户不存在")
	}
	password := sha.EnCode(u.Password)
	if user.Password == password {
		str, _ := middleware.CreateToken(user)
		c.JSON(http.StatusOK, utils.JSONSuccess(
			"",
			str,
		))
		return
	}
	panic("密码错误")
}

func initAdmin(c *gin.Context) {

	ip := c.ClientIP()
	ip = strings.Split(ip, ":")[0]
	if ip != "127.0.0.1" {
		panic(ip)
	}

	secret := c.Query("secret")
	if secret != "fqEeEPlgFECywkwqVMoCEmBzmRmFPZwt" {
		panic("密钥错误")
	}
	db := model.DB

	a := &model.User{Name: "admin", IsAdmin: 1}
	if findAdmin := db.Where(a).Find(a).RowsAffected; findAdmin >= 1 {
		a.Password = sha.AesDecryptCFB(a.Password)
		c.JSON(http.StatusOK, utils.JSONSuccess("", a))
		return
	}
	pwd := tools.RandStringBytes(16)

	admin := &model.User{
		Name:     "admin",
		Password: sha.EnCode(pwd),
		IsAdmin:  1,
	}

	if err := db.Save(admin).Error; err != nil {
		panic(err)
	}
	admin.Password = sha.AesDecryptCFB(admin.Password)
	c.JSON(http.StatusOK, utils.JSONSuccess("", admin))
}
