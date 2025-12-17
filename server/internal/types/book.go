package types

type Book struct {
	Title         string   `json:"title"`
	Author        []string `json:"authors"`
	Description   string   `json:"description"`
	Publisher     string   `json:"publisher"`
	PublishedDate string   `json:"publishedDate"`
	PageCount     int      `json:"pageCount"`
	Language      string   `json:"language"`
}
