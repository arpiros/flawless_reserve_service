package controllers

import (
	"reserve_service/models"

	"github.com/labstack/echo/v4"
)

func GetMember(c echo.Context) error {
	return nil
}

func PutMember(c echo.Context) error {
	member := models.Member{}
	if err := c.Bind(&member); err != nil {
		return err
	}
	if err := c.Validate(&member); err != nil {
		return err
	}

	return nil
}

func DeleteMember(c echo.Context) error {
	return nil
}
