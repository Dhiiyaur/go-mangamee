package controller

import (
	"fmt"
	"net/http"

	"github.com/dhiiyaur/go-mangamee/internal/models"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {

	var user models.User
	if err := c.Bind(&user); err != nil {

		return err
	}

	fmt.Println(user)
	return c.JSON(http.StatusOK, user)

}

func Register(c echo.Context) error {

	var user models.User
	if err := c.Bind(user); err != nil {

		return err
	}

	fmt.Println(user)
	return c.JSON(http.StatusOK, user)

}
