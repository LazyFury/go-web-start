package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/lazyfury/go-web-start/server/model"
	"github.com/lazyfury/go-web-template/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func handleErr(c *gin.Context, err string) {
	if response.IsReqFromHTML(c) {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.AbortWithStatusJSON(http.StatusForbidden, response.JSON(response.AuthedError, err, nil))
}

// Auth 必须登录 后台接口
var Auth gin.HandlerFunc = func(c *gin.Context) {
	log.Printf("auth middleware")
	token, err := getToken(c)
	if err != nil || token == "" {
		handleErr(c, "")
		return
	}
	user, err := parseToken(token)
	if err != nil {
		handleErr(c, "解析token错误")
		return
	}
	c.Set("user", user)
}

// AuthOrNot 兼容前台接口 公开和用户的数据
var AuthOrNot gin.HandlerFunc = func(c *gin.Context) {
	token, err := getToken(c)
	if err != nil || token == "" {
		return
	}
	user, err := parseToken(token)
	if err != nil {
		return
	}
	c.Set("user", user)
}

const (
	// SECRET SECRET
	SECRET string = "asdhjsdhhdhdhdhsasd"
)

func getToken(c *gin.Context) (token string, err error) {
	// query
	token = c.Query("token")
	req := c.Request
	if token != "" {
		return
	}

	// post
	token = c.PostForm("token")
	if token != "" {
		return
	}

	token = req.FormValue("token")
	if token != "" {
		return
	}

	// header
	token = req.Header.Get("token")
	if token != "" {
		return
	}

	token = req.Header.Get("Authorization")
	if token != "" {
		return
	}

	// cookie
	token, err = c.Cookie("token")
	if err != nil {
		return
	}

	// post json token内不做了，需要拷贝一份body，对性能有一些影响

	return
}

// CreateToken 创建token
func CreateToken(u model.User) (token string, err error) {
	return CreateTokenMaxAge(u, int64(60*60*24))
}

// CreateTokenMaxAge 创建token
func CreateTokenMaxAge(u model.User, maxAge int64) (tokens string, err error) {
	//自定义claim
	claim := jwt.MapClaims{
		"id":      u.ID,
		"nick":    u.Name,
		"headPic": "",
		"nbf":     time.Now().Unix(),          //指定时间之前 token不可用
		"iat":     time.Now().Unix(),          //签发时间
		"exp":     time.Now().Unix() + maxAge, //过期时间 24小时
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokens, err = token.SignedString([]byte(SECRET))
	return
}

// 解密token方法
func secret() jwt.Keyfunc {
	key := []byte(SECRET)
	return func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}
}

//ParseToken 解密token
func parseToken(tokens string) (user *model.User, err error) {
	token, err := jwt.Parse(tokens, secret())
	if err != nil {
		err = errors.New("解析token出错")
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to map claim")
		return
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}
	user = &model.User{}
	user.ID = uint(claim["id"].(float64)) // uint64(claim["id"].(float64))
	user.Name = claim["nick"].(string)
	// user.HeadPic = claim["headPic"].(string)

	exp := int64(claim["exp"].(float64))
	fmt.Println(user.Name, "过期时间=====", time.Unix(exp, 0).Format("2001-01-02 15:04:05"))
	return
}
