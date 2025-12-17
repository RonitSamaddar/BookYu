package handlers

import (
	"github.com/RonitSamaddar/BookYu/internal/types"
)

type Handler struct {
	source Source
}

type Source interface {
	GetBook(name string) (types.Book, error)
}
