package sys

import (
	"reserve_service/controllers"

	"github.com/labstack/echo/v4"
)

func SetRouter(e *echo.Echo) {
	e.GET("/member", controllers.GetMember)
	e.PUT("/member", controllers.PutMember)
	e.DELETE("/member/:id", controllers.DeleteMember)
}
