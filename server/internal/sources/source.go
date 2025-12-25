package sources

import (
	"github.com/RonitSamaddar/BookYu/internal/consts"
	gbook "github.com/RonitSamaddar/BookYu/internal/sources/googlebooks"
	"github.com/RonitSamaddar/BookYu/internal/types"
)

func New(config types.Config) *Source {
	return &Source{
		name:   consts.SourceNameGoogleBooks,
		config: config,
	}
}

func (s *Source) GetBook(name string, language string) (types.Book, error) {
	query := gbook.ProcessNameQuery(name)
	return gbook.GetBook(query, language, s.config.GoogleBooksApiKey, 1)
}
