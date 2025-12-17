package types

import "time"

type Book struct {
	Id     string `bson:"_id"`
	Isbn10 string `bson:"isbn10"`
	Isbn13 string `bson:"isbn13"`

	Title         string   `json:"title" bson:"title"`
	Subtitle      string   `json:"subtitle" bson:"subtitle"`
	Authors       []string `json:"authors" bson:"authors"`
	Publishers    string   `json:"publisher" bson:"publisher"`
	PublishedYear string   `bson:"publishedDate"`
	Description   string   `json:"description" bson:"description"`
	Genres        []string `bson:"genres"`
	PageCount     int      `json:"pageCount" bson:"pageCount"`

	AverageRating float64 `json:"averageRating" bson:"averageRating"`
	RatingsCount  int     `json:"ratingsCount" bson:"ratingsCount"`
	Thumbnail     string  `json:"thumbnail" bson:"thumbnail"`

	Source    string    `bson:"source"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

type BookResponse struct {
	Isbn10        string   `json:"isbn10"`
	Isbn13        string   `json:"isbn13"`
	Title         string   `json:"title"`
	Subtitle      string   `json:"subtitle,omitempty"`
	Authors       []string `json:"authors,omitempty"`
	Publishers    string   `json:"publisher,omitempty"`
	PublishedYear string   `bson:"publishedDate,omitempty"`
	Description   string   `json:"description,omitempty"`
	Genres        []string `bson:"genres,omitempty"`
	PageCount     int      `json:"pageCount,omitempty"`
	AverageRating float64  `json:"averageRating,omitempty"`
	RatingsCount  int      `json:"ratingsCount,omitempty"`
	Thumbnail     string   `json:"thumbnail,omitempty"`
}
