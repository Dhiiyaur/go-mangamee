package controller

import (
	"net/http"
	"strings"

	"github.com/dhiiyaur/go-mangamee/internal/service"
	"github.com/labstack/echo/v4"
)

func Browse(c echo.Context) error {

	params := c.QueryParam("pageNumber")
	data := service.BrowsePopularManga(params)

	return c.JSON(http.StatusOK, data)
}

func Search(c echo.Context) error {

	title := strings.Replace(c.QueryParam("mangaTitle"), " ", "%20", -1)
	lang := c.QueryParam("lang")

	if lang == "EN" {

		data, err := service.EnMangaName(title)
		if err != nil {
			// log.Fatal(err)
			return c.JSON(http.StatusBadRequest, err)

		}

		return c.JSON(http.StatusOK, data)

	} else {

		data, err := service.IDMangaName(title)
		if err != nil {
			// log.Fatal(err)
			return c.JSON(http.StatusBadRequest, err)

		}

		return c.JSON(http.StatusOK, data)
	}
}

func Manga(c echo.Context) error {

	title := c.QueryParam("mangaTitle")
	lang := c.QueryParam("lang")

	if lang == "EN" {

		data, err := service.EnMangaChapter(title)
		if err != nil {
			// log.Fatal(err)
			return c.JSON(http.StatusBadRequest, err)

		}

		return c.JSON(http.StatusOK, data)

	} else {
		data, err := service.IDMangaChapter(title)
		if err != nil {
			// log.Fatal(err)
			return c.JSON(http.StatusBadRequest, err)

		}

		return c.JSON(http.StatusOK, data)

	}
}

func Page(c echo.Context) error {

	lang := c.QueryParam("lang")
	title := c.QueryParam("mangaTitle")
	chapter := c.QueryParam("chapter")

	if lang == "EN" {

		data := service.EnMangaImage(title, chapter)
		return c.JSON(http.StatusOK, data)

	} else {

		data := service.IDMangaImage(chapter)
		return c.JSON(http.StatusOK, data)
	}
}
