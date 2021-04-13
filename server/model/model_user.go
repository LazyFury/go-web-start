package model

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lazyfury/go-web-start/server/config"
	"github.com/lazyfury/go-web-template/response"
	"github.com/lazyfury/go-web-template/tools/types"
	"github.com/lazyfury/go-web-template/tools/wechat"
)

// UserID UserID
type UserID struct {
	UserID uint `json:"user_id" gorm:"not null"`
}

// User 用户更新
type User struct {
	BaseControll
	Password  string          `json:"password" gorm:"not null;type:text"`
	Name      string          `json:"name" gorm:"unique;not null"`
	Email     string          `json:"email"`
	IP        string          `json:"ip"`
	Ua        string          `json:"ua"`
	LoginTime types.LocalTime `json:"login_time"`
	Status    int             `json:"status"`
	IsAdmin   int             `json:"is_admin" gorm:"default:0"`
}

// WechatOauth 微信用户登陆
type WechatOauth struct {
	BaseControll
	UID int `json:"user_id"`
	wechat.UserInfo
}

// PointerList 列表
func (u *User) PointerList() interface{} {
	type tmp struct {
		*User
		Password string `json:"-"`
	}
	return &[]tmp{}
}

// Pointer 实例
func (u *User) Pointer() interface{} {
	return &User{}
}

// TableName 表名
func (u *User) TableName() string {
	return TableName("users")
}

// Add 添加
func (u *User) Add(c *gin.Context) {
	user := &User{}

	if err := c.Bind(user); err != nil {
		response.Error(response.JSONError("参数错误", err))
	}

	user.Name = strings.Trim(user.Name, " ")
	if user.Name == "" {
		response.Error(response.JSONError("用户名不可空", nil))
	}
	if user.Password == "" {
		response.Error(response.JSONError("用户密码不可空", nil))
	}

	user.Password = config.Global.Sha1.EnCode(user.Password)

	req := c.Request
	ua := req.UserAgent()
	ip := c.ClientIP()
	user.IP = ip
	user.Ua = ua
	user.LoginTime = types.LocalTime{Time: time.Now()}
	user.Status = 1

}

// Update 更新
func (u *User) Update(c *gin.Context) {
	user := new(User)

	if err := c.Bind(user); err != nil {
		response.Error(response.JSONError("参数错误", err))
	}

	_u := &User{BaseControll: BaseControll{ID: uint(user.ID)}}
	err := _u.HasUser()
	if err != nil {
		response.Error(err)
	}

	user.Name = strings.Trim(user.Name, " ")
	if user.Password != "" {
		user.Password = config.Global.Sha1.EnCode(user.Password)
	}

}

// DelUser 删除用户
func (u *User) DelUser() (interface{}, error) {
	db := DB
	row := db.Delete(u)
	// row := db.Exec("DELETE FROM "+config.Global.TablePrefix+"`users` WHERE `id` = ?", u.ID)
	if row.Error != nil {
		return row.Error, errors.New("删除失败")
	}
	if row.RowsAffected <= 0 {
		return nil, errors.New("删除失败,数据不存在")
	}
	return nil, nil
}

// Frozen 冻结用户
func (u *User) Frozen(c *gin.Context) {
	user := new(User)

	if err := c.Bind(&user); err != nil {
		response.Error(err)
	}

	db := DB
	row := db.Model(&User{BaseControll: BaseControll{ID: user.ID}}).Update("status", user.Status)
	if row.Error != nil {
		response.Error("操作失败")
	}

	if user.Status == 0 {
		response.Error("冻结用户")
	}

	c.JSON(http.StatusOK, response.JSONSuccess("解冻用户", nil))
}

// HasUser 查找用户
func (u *User) HasUser() error {
	db := DB
	row := db.Where(u).First(u)
	if row.Error != nil {
		return errors.New("用户不存在")
	}
	return nil
}

// RepeatOfEmail RepeatOfEmail
func (u *User) RepeatOfEmail(c *gin.Context) {
	user := new(User)
	email := c.Query("email")
	user.Email = email
	err := user.HasUser()
	if err != nil {
		c.JSON(http.StatusOK, response.JSONSuccess("没有重复", nil))
		return
	}
	c.JSON(http.StatusOK, response.JSON(response.RepeatEmail, "", nil))
}

// RepeatOfName RepeatOfName
func (u *User) RepeatOfName(c *gin.Context) {
	user := new(User)
	name := c.Query("name")
	user.Name = name
	err := user.HasUser()
	if err != nil {
		response.Error("没有重复")
	}
	c.JSON(http.StatusOK, response.JSON(response.RepeatUserName, "", nil))
}
