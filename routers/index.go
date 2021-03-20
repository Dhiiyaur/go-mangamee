package routers

import (
	"net/http"

	"github.com/dhiiyaur/go-mangamee/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Index() *echo.Echo {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "MangameeApi")
	})

	e.GET("/browse", func(c echo.Context) error {

		params := c.QueryParam("page")
		data := models.BrowsePopularManga(params)
		return c.JSON(http.StatusOK, data)
	})

	e.GET("/search", func(c echo.Context) error {

		title := c.QueryParam("mangaTitle")
		lang := c.QueryParam("lang")
		if lang == "EN" {

			data, err := models.EnMangaName(title)
			if err != nil {
				// log.Fatal(err)
				return c.JSON(http.StatusBadRequest, err)

			}

			return c.JSON(http.StatusOK, data)

		} else {

			data, err := models.IDMangaName(title)
			if err != nil {
				// log.Fatal(err)
				return c.JSON(http.StatusBadRequest, err)

			}

			return c.JSON(http.StatusOK, data)
		}

	})

	e.GET("/manga", func(c echo.Context) error {

		title := c.QueryParam("mangaTitle")
		lang := c.QueryParam("lang")

		if lang == "EN" {

			data, err := models.EnMangaChapter(title)
			if err != nil {
				// log.Fatal(err)
				return c.JSON(http.StatusBadRequest, err)

			}

			return c.JSON(http.StatusOK, data)

		} else {
			data, err := models.IDMangaChapter(title)
			if err != nil {
				// log.Fatal(err)
				return c.JSON(http.StatusBadRequest, err)

			}

			return c.JSON(http.StatusOK, data)

		}

	})

	e.GET("/page", func(c echo.Context) error {

		lang := c.QueryParam("lang")
		title := c.QueryParam("mangaTitle")
		chapter := c.QueryParam("chapter")

		if lang == "EN" {
			data := models.EnMangaImage(title, chapter)
			return c.JSON(http.StatusOK, data)
		} else {
			data := models.IDMangaImage(chapter)
			return c.JSON(http.StatusOK, data)
		}

	})

	return e
}
