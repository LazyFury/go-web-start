package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getCsrfToken(c echo.Context) error {
	return c.Blob(http.StatusOK, "text/html;charset=utf-8;", []byte(""))
}
