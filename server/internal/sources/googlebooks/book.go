package googlebooks

import (
	"fmt"
	"time"

	"github.com/RonitSamaddar/BookYu/internal/consts"
	"github.com/RonitSamaddar/BookYu/internal/types"
)

func processVolumeInfo(volumeInfo GoogleBooksVolumeInfo) (types.Book, error) {
	book := types.Book{
		Title:         volumeInfo.Title,
		Subtitle:      volumeInfo.Subtitle,
		Authors:       volumeInfo.Authors,
		Publishers:    volumeInfo.Publisher,
		PublishedYear: volumeInfo.PublishedDate,
		Description:   volumeInfo.Description,
		Genres:        volumeInfo.Categories,
		PageCount:     volumeInfo.PageCount,
		AverageRating: volumeInfo.AverageRating,
		RatingsCount:  volumeInfo.RatingsCount,
		Thumbnail:     volumeInfo.ImageLinks.Thumbnail,
		Source:        consts.SourceNameGoogleBooks,
		UpdatedAt:     time.Now(),
	}

	for _, volume := range volumeInfo.IndustryIdentifiers {
		switch volume.Type {
		case consts.GoogleBooksIsbn13Type:
			book.Isbn13 = volume.Identifier
		case consts.GoogleBooksIsbn10Type:
			book.Isbn10 = volume.Identifier
		}
	}

	if len(book.Isbn13) == 0 && len(book.Isbn10) == 0 {
		return types.Book{}, fmt.Errorf("no isbn code found")
	} else if len(book.Isbn13) != 0 {
		book.Id = fmt.Sprintf("%s-%s", book.Title, book.Isbn13)
	} else {
		book.Id = fmt.Sprintf("%s-%s", book.Title, book.Isbn10)
	}

	return book, nil
}
