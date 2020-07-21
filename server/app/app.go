package app

import (
	"github.com/labstack/echo"
)

// App App
type App struct {
	*echo.Echo
}

// Get Get
func (a *App) Get(path, callback func(c echo.Context) error) error {
	return nil
}
