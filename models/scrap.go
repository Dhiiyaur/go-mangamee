package models

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// EN ------------------------------------------------------------------------------------------

func BrowsePopularManga(page string) interface{} {

	type Manga struct {
		MangaCover string `json:"MangaCover"`
		MangaLink  string `json:"MangaLink"`
		MangaTitle string `json:"MangaTitle"`
	}

	type Mangas []Manga

	var DataMangas Mangas

	c := colly.NewCollector()

	c.OnHTML(".media-left", func(e *colly.HTMLElement) {

		MangaCover := e.ChildAttr("a img", "src")
		MangaLink := strings.Split(e.ChildAttr(`a`, "href"), "/")[4]
		MangaTitle := strings.Split(strings.Replace(strings.Split(e.ChildAttr(`a`, "href"), "/")[4], "-", " ", -1), "_")[0]

		result := Manga{

			MangaCover: MangaCover,
			MangaLink:  MangaLink,
			MangaTitle: MangaTitle,
		}

		DataMangas = append(DataMangas, result)

	})

	c.Visit("https://mangahub.io/popular/page/" + page + "/")

	return DataMangas

}

func EnMangaName(title string) (interface{}, error) {

	type Manga struct {
		MangaCover   string `json:"MangaCover"`
		MangaLink    string `json:"MangaLink"`
		MangaTitle   string `json:"MangaTitle"`
		MangaChapter string `json:"MangaChapter"`
		MangaStatus  string `json:"MangaStatus"`
	}

	type Mangas []Manga

	var DataMangas Mangas

	c := colly.NewCollector()

	c.OnHTML(".media-manga", func(h *colly.HTMLElement) {

		MangaCover := "temp"

		h.ForEach(".media-left", func(i int, h *colly.HTMLElement) {

			MangaCover = h.ChildAttr("a img", "src")

		})

		h.ForEach(".media-body", func(i int, h *colly.HTMLElement) {

			MangaLink := strings.Split(h.ChildAttr(`a`, "href"), "/")[4]
			MangaTitle := h.ChildText("h4 a")
			MangaChapter := strings.Split(h.ChildText("p"), "published")[0]
			MangaStatus := strings.Split(strings.Split(h.ChildText("p"), "(")[1], ")")[0]

			result := Manga{

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

		return "", errors.New("empty")
	}

	return DataMangas, nil
}

func EnMangaChapter(title string) (interface{}, error) {

	type Chapter struct {
		ChapterLink string `json:"ChapterLink"`
		ChapterName string `json:"ChapterName"`
	}

	type Chapters []Chapter

	type MangaData struct {
		CoverImage string   `json:"CoverImage"`
		Summary    string   `json:"Summary"`
		Chapter    Chapters `json:"Chapter"`
	}
	var MangaChapter Chapters

	var coverImage, summary string

	c := colly.NewCollector()
	c.OnHTML("._2U6DJ", func(e *colly.HTMLElement) {

		ChapterLink := strings.Split(e.Attr("href"), "/")[5]
		ChapterName := strings.Split(e.Attr("href"), "/")[5]

		result := Chapter{

			ChapterName: ChapterName,
			ChapterLink: ChapterLink,
		}
		MangaChapter = append(MangaChapter, result)

	})

	c.OnHTML("._4RcEi", func(e *colly.HTMLElement) {

		coverImage = e.ChildAttr("img", "src")

	})

	c.OnHTML(".ZyMp7", func(e *colly.HTMLElement) {

		summary = e.Text

	})

	c.Visit("https://mangahub.io/manga/" + title)

	data := MangaData{

		CoverImage: coverImage,
		Summary:    summary,
		Chapter:    MangaChapter,
	}

	if len(MangaChapter) == 0 {

		return "", errors.New("empty")

	}
	return data, nil

}

func EnMangaImage(title string, chapter string) interface{} {

	type Image struct {
		Image string `json:"Image"`
	}

	type Images []Image

	type MangaData struct {
		Image Images `json:"Image"`
	}

	var link, frond, end string
	var MangaImage Images

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
		tempImage := Image{

			Image: data,
		}

		MangaImage = append(MangaImage, tempImage)

	}

	data := MangaData{

		Image: MangaImage,
	}

	return data
}

// IND ------------------------------------------------------------------------------------------

func IDMangaName(title string) (interface{}, error) {

	type Manga struct {
		MangaCover   string `json:"MangaCover"`
		MangaLink    string `json:"MangaLink"`
		MangaTitle   string `json:"MangaTitle"`
		MangaChapter string `json:"MangaChapter"`
		MangaStatus  string `json:"MangaStatus"`
	}

	type Mangas []Manga

	var DataMangas Mangas

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

			result := Manga{

				MangaCover:   MangaCover,
				MangaLink:    MangaLink,
				MangaTitle:   MangaTitle,
				MangaChapter: MangaChapter,
				MangaStatus:  "-",
			}

			DataMangas = append(DataMangas, result)

		})

	})

	c.Visit("https://www.maid.my.id/?s=" + title)

	if len(DataMangas) == 0 {

		return "", errors.New("empty")
	}

	return DataMangas, nil
}

func IDMangaChapter(title string) (interface{}, error) {

	type Chapter struct {
		ChapterLink string `json:"ChapterLink"`
		ChapterName string `json:"ChapterName"`
	}

	type Chapters []Chapter

	type MangaData struct {
		CoverImage string   `json:"CoverImage"`
		Summary    string   `json:"Summary"`
		Chapter    Chapters `json:"Chapter"`
	}

	var MangaChapter Chapters

	var coverImage, summary string

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

		result := Chapter{

			ChapterName: ChapterName,
			ChapterLink: ChapterLink,
		}
		MangaChapter = append(MangaChapter, result)

	})
	c.Visit("https://www.maid.my.id/manga/" + title + "/")

	data := MangaData{

		CoverImage: coverImage,
		Summary:    summary,
		Chapter:    MangaChapter,
	}

	if len(MangaChapter) == 0 {

		return "", errors.New("empty")

	}
	return data, nil
}

func IDMangaImage(chapter string) interface{} {

	type Image struct {
		Image string `json:"Image"`
	}

	type Images []Image

	type MangaData struct {
		Image Images `json:"Image"`
	}

	var MangaImage Images

	c := colly.NewCollector()
	c.OnHTML(".reader-area img", func(e *colly.HTMLElement) {

		data := e.Attr("src")
		tempImage := Image{

			Image: data,
		}

		MangaImage = append(MangaImage, tempImage)

	})

	c.Visit("https://www.maid.my.id/" + chapter + "/")

	data := MangaData{
		Image: MangaImage,
	}

	return data
}
