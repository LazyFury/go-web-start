package model

import (
	"EK-Server/util"
	"EK-Server/util/customtype"
	"EK-Server/util/sha"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// User 用户更新
type User struct {
	gorm.Model
	Password  string               `json:"password" gorm:"not null"`
	Name      string               `json:"name" gorm:"unique;not null"`
	Email     string               `json:"email"`
	IP        string               `json:"ip"`
	Ua        string               `json:"ua"`
	LoginTime customtype.LocalTime `json:"login_time"`
	Status    int                  `json:"status"`
	IsAdmin   bool                 `json:"isAdmin" gorm:"default:0"`
}

// SearchUser	 用户列表显示
type searchUser struct {
	ID        uint                 `json:"id"`
	Email     string               `json:"email"`
	Name      string               `json:"name"`
	IP        string               `json:"ip"`
	Ua        string               `json:"ua"`
	LoginTime customtype.LocalTime `json:"login_time"`
	Status    int                  `json:"status"`
}

// WechatOauth 微信用户登陆
type WechatOauth struct {
	gorm.Model
	UID          int    `json:"uid"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Nickname     string `json:"nickname"`
	Sex          int    `json:"sex"`
	Headimgurl   string `json:"headimgurl"`
	Province     string `json:"province"`
	City         string `json:"city"`
	Country      string `json:"country"`
	Unionid      string `json:"unionid"`
	Subscribe    int    `json:"subscribe"`
}

// Find 查找用户
func (u *User) Find() error {
	db := DB
	if db.Where(u).First(u).RecordNotFound() {
		return errors.New("用户不存在")
	}
	return nil
}

// GetAllUser  获取所有用户列表
func (u *User) GetAllUser(limit int, page int) *Result {
	return DataBaselimit(limit, page, map[string]interface{}{}, &[]searchUser{}, "users", "id desc")
}

// RegController AddUser
func (u *User) RegController(c echo.Context) error {

	user := new(User)

	if err := c.Bind(user); err != nil {
		return util.JSONErr(c, err, "参数错误")
	}

	user.Name = strings.Trim(user.Name, " ")
	if user.Name == "" {
		return util.JSONErr(c, nil, "用户名不可空")
	}

	user.Password = sha.EnCode(user.Password)

	req := c.Request()
	ua := req.UserAgent()
	ip := util.ClientIP(c)
	user.IP = ip
	user.Ua = ua
	user.LoginTime = customtype.LocalTime{Time: time.Now()}
	user.Status = 1

	fmt.Println(user)
	msg, err := user.AddUser()
	if err != nil {
		return util.JSONErr(c, nil, msg)
	}
	return util.JSONSuccess(c, user.ID, msg)
}

// UpdateUser 更新用户
func (u *User) UpdateUser(id uint, data *User) error {
	db := DB

	// uid, _ := strconv.Atoi(id)
	// 使用id查找用户
	user := &User{Model: gorm.Model{ID: uint(id)}}
	err := user.Find()
	if err != nil {
		return err
	}
	row := db.Model(user).Updates(data)
	if row.Error != nil {
		return row.Error
	}
	if row.RowsAffected <= 0 {
		return errors.New("没有更改")
	}
	return nil
}

// AddUser 添加用户
func (u *User) AddUser() (string, error) {
	db := DB

	// fmt.Println(u)
	db.NewRecord(u) // => 主键为空返回`true`
	row := db.Create(u)

	if row.Error != nil {
		return "添加失败,用户已存在", row.Error
	}

	if row.RowsAffected <= 0 {
		return "添加失败，没有更改", errors.New("")
	}

	return "添加成功", nil
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
