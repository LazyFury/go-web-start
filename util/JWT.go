package util

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// UserInfo jwt UserInfo type
type UserInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

const (
	// SECRET jwt
	SECRET string = "token.secret"
)

// AdminJWT 管理后台用户验证
var AdminJWT echo.MiddlewareFunc = baseJWT(adminCheckToken)

// CheckToken 检查token可用
func adminCheckToken(next echo.HandlerFunc, c echo.Context, tokenString string) error {
	user, err := parseToken(tokenString)
	if err != nil {
		return JSON(c, err, "登陆失效!", -101)
	}
	if user.ID == "" {
		return JSON(c, nil, "登陆失效!", -101)
	}
	c.Request().Header.Add("uid", user.ID)
	return next(c)
}

// UserJWT 普通用户验证
var UserJWT echo.MiddlewareFunc = baseJWT(userCheckToken)

// CheckToken 检查token可用
func userCheckToken(next echo.HandlerFunc, c echo.Context, token string) error {
	if token != "312" {
		return JSON(c, nil, "登陆失效!", -101)
	}
	return next(c)
}

// JWT 验证
func baseJWT(callback func(next echo.HandlerFunc, c echo.Context, token string) error) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		fmt.Printf("\n>>>>>>>Base Check>>>>>>>>>\n")
		return func(c echo.Context) error {
			var token string = getToken(c)

			if token != "" {
				return callback(next, c, token)
			}

			return JSON(c, nil, "请先登陆!", -100)
		}
	}
}

// 获取token
func getToken(c echo.Context) (token string) {

	token = c.QueryParam("token")
	if token != "" {
		return token
	}

	r := c.Request()
	token = r.Header.Get("token")
	if token != "" {
		return token
	}
	token = r.Header.Get("Authorization")
	if token != "" {
		return token
	}
	return ""
}

// CreateToken 创建token
func CreateToken(user *UserInfo) (tokenstr string, err error) {
	//自定义claim
	claim := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Name,
		"nbf":      time.Now().Unix(),            //指定时间之前 token不可用
		"iat":      time.Now().Unix(),            //签发时间
		"exp":      time.Now().Unix() + 60*60*24, //过期时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenstr, err = token.SignedString([]byte(SECRET))
	fmt.Printf(tokenstr)
	return
}

// 解密token方法
func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	}
}

//ParseToken 解密token
func parseToken(tokenss string) (user *UserInfo, err error) {
	user = &UserInfo{}
	token, err := jwt.Parse(tokenss, secret())
	if err != nil {
		err = errors.New("解析token出错")
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}

	user.ID = claim["id"].(string) // uint64(claim["id"].(float64))
	user.Name = claim["username"].(string)

	exp := int64(claim["exp"].(float64))
	fmt.Println(user, "过期时间=====", time.Unix(exp, 0).Format(timeLayout))
	return
}
