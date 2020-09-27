package middleware

import "github.com/labstack/echo/v4"

func SetMiddleWare(e *echo.Echo) {
	e.Validator = NewValidator()
	e.Use(NewRecover())
	e.Use(HttpLogger)
}
