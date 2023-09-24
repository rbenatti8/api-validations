package handler

import "github.com/labstack/echo/v4"

type UserHandler struct {
}

func (u *UserHandler) Create(c echo.Context) error {
	var user CreateUserDTO
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	return c.JSON(200, user)
}
