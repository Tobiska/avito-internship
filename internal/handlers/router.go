package handlers

import (
	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	return echo.New()
}
