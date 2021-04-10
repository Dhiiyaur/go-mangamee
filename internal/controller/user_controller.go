package controller

import (
	"fmt"
	"net/http"

	"github.com/dhiiyaur/go-mangamee/internal/models"
	"github.com/dhiiyaur/go-mangamee/internal/service"
	"github.com/dhiiyaur/go-mangamee/internal/utils"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {

	var user models.User
	if err := c.Bind(&user); err != nil {

		return err
	}

	result, err := service.Login(user)

	if err != nil {

		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func Register(c echo.Context) error {

	var user models.User
	if err := c.Bind(&user); err != nil {

		return err
	}

	result, err := service.Register(user)

	if err != nil {

		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func GetHistory(c echo.Context) error {

	token := c.Request().Header.Get("Authorization")

	if token == "" {
		return echo.ErrUnauthorized
	}

	username, _ := utils.CheckToken(token)

	result, err := service.GetHistory(username)

	if err != nil {

		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// down

func DeleteHistory(c echo.Context) error {

	token := c.Request().Header.Get("Authorization")

	if token == "" {
		return echo.ErrUnauthorized
	}

	username, _ := utils.CheckToken(token)

	var userHistory models.UserHistory
	if err := c.Bind(&userHistory); err != nil {

		return err
	}

	err := service.DeleteHistory(username, userHistory)

	if err != nil {

		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "delete ok")
}

func CreateHistory(c echo.Context) error {

	token := c.Request().Header.Get("Authorization")

	if token == "" {
		return echo.ErrUnauthorized
	}

	username, _ := utils.CheckToken(token)

	var userHistory models.UserHistory
	if err := c.Bind(&userHistory); err != nil {

		return err
	}

	err := service.CreateHistory(username, userHistory)

	if err != nil {

		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "created ok")
}
