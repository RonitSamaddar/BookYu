package server

import (
	"github.com/RonitSamaddar/BookYu/internal/types"
)

type Server struct {
	config  *types.Config
	handler Handler
}

type Handler interface {
	RegisterRoutes()
}
