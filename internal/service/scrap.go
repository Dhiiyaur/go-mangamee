package service

import (
	"errors"
	"strconv"
	"strings"

	"github.com/dhiiyaur/go-mangamee/internal/models"
	"github.com/gocolly/colly"
)

// EN ------------------------------------------------------------------------------------------

func BrowsePopularManga(page string) ([]models.MangaData, error) {

	DataMangas := []models.MangaData{}

	c := colly.NewCollector()
	c.OnHTML(".media-left", func(e *colly.HTMLElement) {

		MangaCover := e.ChildAttr("a img", "src")
		MangaLink := strings.Split(e.ChildAttr(`a`, "href"), "/")[4]
		MangaTitle := strings.Split(strings.Replace(strings.Split(e.ChildAttr(`a`, "href"), "/")[4], "-", " ", -1), "_")[0]

		result := models.MangaData{

			MangaCover: MangaCover,
			MangaLink:  MangaLink,
			MangaTitle: MangaTitle,
		}

		DataMangas = append(DataMangas, result)

	})

	c.Visit("https://mangahub.io/popular/page/" + page + "/")

	if len(DataMangas) == 0 {

		return DataMangas, errors.New("empty")
	}

	return DataMangas, nil

}

func EnMangaName(title string) ([]models.MangaData, error) {

	DataMangas := []models.MangaData{}

	c := colly.NewCollector()
	c.OnHTML(".media-manga", func(h *colly.HTMLElement) {

		MangaCover := "temp"

		h.ForEach(".media-left", func(i int, h *colly.HTMLElement) {

			MangaCover = h.ChildAttr("a img", "src")

		})

		h.ForEach(".media-body", func(i int, h *colly.HTMLElement) {

			MangaLink := strings.Split(h.ChildAttr(`a`, "href"), "/")[4]
			MangaTitle := h.ChildText("h4 a")
			MangaChapter := strings.Split(strings.Split(h.ChildText("p"), "published")[0], " ")[0]
			MangaStatus := strings.Split(strings.Split(h.ChildText("p"), "(")[1], ")")[0]

			result := models.MangaData{

				MangaCover:   MangaCover,
				MangaLink:    MangaLink,
				MangaTitle:   MangaTitle,
				MangaChapter: MangaChapter,
				MangaStatus:  MangaStatus,
			}

			DataMangas = append(DataMangas, result)

		})

	})

	c.Visit("https://mangahub.io/search?q=" + title)

	if len(DataMangas) == 0 {

		return DataMangas, errors.New("empty")
	}

	return DataMangas, nil
}

func EnMangaChapter(title string) (models.MangaData, error) {

	mangaChapter := []models.Chapter{}
	var coverImage, summary string

	c := colly.NewCollector()
	c.OnHTML("._2U6DJ", func(e *colly.HTMLElement) {

		ChapterLink := strings.Split(e.Attr("href"), "/")[5]
		ChapterName := strings.Split(e.Attr("href"), "/")[5]

		result := models.Chapter{

			ChapterName: ChapterName,
			ChapterLink: ChapterLink,
		}

		mangaChapter = append(mangaChapter, result)

	})

	c.OnHTML("._4RcEi", func(e *colly.HTMLElement) {

		coverImage = e.ChildAttr("img", "src")

	})

	c.OnHTML(".ZyMp7", func(e *colly.HTMLElement) {

		summary = e.Text

	})

	c.Visit("https://mangahub.io/manga/" + title)

	DataMangas := models.MangaData{

		MangaCover: coverImage,
		Summary:    summary,
		Chapter:    mangaChapter,
	}

	if len(mangaChapter) == 0 {

		return DataMangas, errors.New("empty")

	}

	return DataMangas, nil

}

func EnMangaImage(title string, chapter string) (models.MangaData, error) {

	var link, frond, end string
	mangaImage := []models.Image{}

	c := colly.NewCollector()
	c.OnHTML(".PB0mN", func(e *colly.HTMLElement) {

		link = e.Attr("src")

	})

	c.Visit("https://mangahub.io/chapter/" + title + "/" + chapter)

	temp := strings.Split(link, "/")
	frond = strings.Join(temp[:len(temp)-1], "/")
	end = strings.Split(temp[len(temp)-1], ".")[1]

	for i := 1; i < 150; i++ {

		data := frond + "/" + strconv.Itoa(i) + "." + end
		tempImage := models.Image{

			Image: data,
		}
		mangaImage = append(mangaImage, tempImage)

	}

	DataMangas := models.MangaData{

		Images: mangaImage,
	}

	if len(mangaImage) == 0 {

		return DataMangas, errors.New("empty")

	}

	return DataMangas, nil
}

// IND ------------------------------------------------------------------------------------------

func IDMangaName(title string) ([]models.MangaData, error) {

	DataMangas := []models.MangaData{}

	c := colly.NewCollector()
	c.OnHTML(".flexbox2-content", func(e *colly.HTMLElement) {

		MangaLink := strings.Split(e.ChildAttr(`a`, "href"), "/")[4]
		MangaChapter := "temp"

		e.ForEach(".season", func(i int, h *colly.HTMLElement) {

			MangaChapter = h.Text
		})

		e.ForEach(".flexbox2-thumb", func(i int, h *colly.HTMLElement) {

			MangaCover := h.ChildAttr(`img`, "src")
			MangaTitle := h.ChildAttr(`img`, "title")

			result := models.MangaData{

				MangaCover:   MangaCover,
				MangaLink:    MangaLink,
				MangaTitle:   MangaTitle,
				MangaChapter: MangaChapter,
			}

			DataMangas = append(DataMangas, result)

		})

	})

	c.Visit("https://www.maid.my.id/?s=" + title)

	if len(DataMangas) == 0 {

		return DataMangas, errors.New("empty")
	}

	return DataMangas, nil
}

func IDMangaChapter(title string) (models.MangaData, error) {

	var coverImage, summary string
	mangaChapter := []models.Chapter{}

	c := colly.NewCollector()

	c.OnHTML(".series-thumb", func(e *colly.HTMLElement) {

		coverImage = e.ChildAttr(`img`, "src")

	})
	c.OnHTML(".series-synops", func(e *colly.HTMLElement) {

		summary = e.Text

	})

	c.OnHTML(".flexch-infoz", func(e *colly.HTMLElement) {

		ChapterName := e.ChildAttr(`a`, "title")
		ChapterLink := strings.Split(e.ChildAttr(`a`, "href"), "/")[3]

		result := models.Chapter{

			ChapterName: ChapterName,
			ChapterLink: ChapterLink,
		}
		mangaChapter = append(mangaChapter, result)

	})
	c.Visit("https://www.maid.my.id/manga/" + title + "/")

	DataMangas := models.MangaData{

		MangaCover: coverImage,
		Summary:    summary,
		Chapter:    mangaChapter,
	}

	if len(mangaChapter) == 0 {

		return DataMangas, errors.New("empty")

	}
	return DataMangas, nil
}

func IDMangaImage(chapter string) (models.MangaData, error) {

	mangaImage := []models.Image{}

	c := colly.NewCollector()
	c.OnHTML(".reader-area img", func(e *colly.HTMLElement) {

		data := e.Attr("src")
		tempImage := models.Image{

			Image: data,
		}

		mangaImage = append(mangaImage, tempImage)

	})

	c.Visit("https://www.maid.my.id/" + chapter + "/")

	DataMangas := models.MangaData{

		Images: mangaImage,
	}

	if len(mangaImage) == 0 {

		return DataMangas, errors.New("empty")

	}

	return DataMangas, nil
}
