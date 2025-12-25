package handlers

import (
	"github.com/RonitSamaddar/BookYu/internal/types"
)

type Handler struct {
	source Source
}

type Source interface {
	GetBook(name, language string) (types.Book, error)
}
