package util

import "github.com/labstack/echo"

// CheckErr CheckErr
func CheckErr(err interface{}, c echo.Context, msg string) {
	if err != nil {
		JSONErr(c, err, msg)
		return
	}
}
