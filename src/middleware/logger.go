package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func HttpLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := time.Now()

		latency := time.Since(t)
		status := c.Response().Status
		method := c.Request().Method
		url := c.Request().RequestURI
		c.Logger().SetHeader(" ")
		c.Logger().Debug("| ", status, " | ", latency, " | ", method, " | ", url)

		m := echo.Map{}
		if err := c.Bind(&m); err != nil {
			return err
		}
		c.Logger().Debug(fmt.Sprintf("%v", m))
		return next(c)
	}
}
