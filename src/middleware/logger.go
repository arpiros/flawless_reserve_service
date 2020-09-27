package middleware

import (
	"github.com/labstack/echo/v4"
	md "github.com/labstack/echo/v4/middleware"
)

func NewLogger() echo.MiddlewareFunc {
	return md.LoggerWithConfig(md.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	})
}
