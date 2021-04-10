package app

import (
	"net/http"
	"os"

	"github.com/dhiiyaur/go-mangamee/internal/controller"
	"github.com/dhiiyaur/go-mangamee/internal/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {

	db.InitDb()

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

	e.GET("/user/gethistory", controller.GetHistory)
	e.POST("/user/deletehistory", controller.DeleteHistory)
	e.POST("/user/createhistory", controller.CreateHistory)

	e.Start(":" + os.Getenv("PORT"))
}
