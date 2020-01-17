package util

import (
	"strings"

	"github.com/labstack/echo"
)

// CheckErr CheckErr
func CheckErr(err interface{}, c echo.Context, msg string) {
	if err != nil {
		JSONErr(c, err, msg)
		return
	}
}

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(c echo.Context) string {
	r := c.Request()
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	return strings.TrimSpace(r.RemoteAddr)
}
