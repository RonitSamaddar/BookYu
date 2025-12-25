package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RonitSamaddar/BookYu/internal/config"
	"github.com/RonitSamaddar/BookYu/internal/consts"
	"github.com/RonitSamaddar/BookYu/internal/handlers"
	"github.com/RonitSamaddar/BookYu/internal/sources"
)

func Start() error {
	server := &Server{}

	config, err := config.LoadInputConfigs()
	if err != nil {
		return err
	}
	server.config = config
	log.Printf("Config loaded : %+v", config)

	if err := server.setup(); err != nil {
		return err
	}

	log.Printf("Starting server")
	log.Printf("Listening on port %s", consts.ServerPort)
	err = http.ListenAndServe(fmt.Sprintf(":%s", consts.ServerPort), nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) setup() error {
	log.Printf("Setting up server")

	source := sources.New(*s.config)

	handler := handlers.New(source)
	s.handler = handler

	handler.RegisterRoutes()
	return nil
}
