package googlebooks

type GoogleBooksGetVolumeResponse struct {
	Kind       string                  `json:"kind"`
	TotalItems int                     `json:"totalItems"`
	Items      []GoogleBooksVolumeItem `json:"items"`
}

type GoogleBooksVolumeItem struct {
	Kind       string                `json:"kind"`
	Id         string                `json:"id"`
	SelfLink   string                `json:"selfLink"`
	VolumeInfo GoogleBooksVolumeInfo `json:"volumeInfo"`
}

type GoogleBooksIndustryIdentifier struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

type GoogleBooksVolumeInfo struct {
	Title         string   `json:"title"`
	Subtitle      string   `json:"subtitle,omitempty"`
	Authors       []string `json:"authors,omitempty"`
	Publisher     string   `json:"publisher,omitempty"`
	PublishedDate string   `json:"publishedDate,omitempty"`
	Description   string   `json:"description,omitempty"`

	// Google key is categories
	Categories []string `json:"categories,omitempty"`

	PageCount     int     `json:"pageCount,omitempty"`
	AverageRating float64 `json:"averageRating,omitempty"`
	RatingsCount  int     `json:"ratingsCount,omitempty"`

	IndustryIdentifiers []GoogleBooksIndustryIdentifier `json:"industryIdentifiers,omitempty"`

	ImageLinks struct {
		SmallThumbnail string `json:"smallThumbnail,omitempty"`
		Thumbnail      string `json:"thumbnail,omitempty"`
	} `json:"imageLinks,omitempty"`

	Language string `json:"language,omitempty"`
}
