package bookManager

import (
	"github.com/RonitSamaddar/BookYu/internal/types"
)

func GetBookResponse(book types.Book) types.BookResponse {
	return types.BookResponse{
		Isbn10:        book.Isbn10,
		Isbn13:        book.Isbn13,
		Title:         book.Title,
		Subtitle:      book.Subtitle,
		Authors:       book.Authors,
		Publishers:    book.Publishers,
		PublishedYear: book.PublishedYear,
		Description:   book.Description,
		Genres:        book.Genres,
		PageCount:     book.PageCount,
		AverageRating: book.AverageRating,
		RatingsCount:  book.RatingsCount,
		Thumbnail:     book.Thumbnail,
	}
}
