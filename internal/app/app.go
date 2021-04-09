package app

import (
	"net/http"
	"os"

	"github.com/dhiiyaur/go-mangamee/internal/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome To MangameeApi")
	})

	e.GET("/browse", controller.Browse)
	e.GET("/search", controller.Search)
	e.GET("/manga", controller.Manga)
	e.GET("/page", controller.Page)

	e.POST("/user/login", controller.Login)
	e.POST("/user/register", controller.Register)

	e.Start(":" + os.Getenv("PORT"))
}
