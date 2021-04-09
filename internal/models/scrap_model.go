package models

type MangaData struct {
	MangaCover   string    `json:"MangaCover"`
	MangaLink    string    `json:"MangaLink"`
	MangaTitle   string    `json:"MangaTitle"`
	MangaChapter string    `json:"MangaChapter"`
	MangaStatus  string    `json:"MangaStatus"`
	Summary      string    `json:"Summary"`
	Chapter      []Chapter `json:"Chapter"`
	Images       []Image   `json:"Image"`
}

type Chapter struct {
	ChapterLink string `json:"ChapterLink"`
	ChapterName string `json:"ChapterName"`
}

type Image struct {
	Image string `json:"Image"`
}
