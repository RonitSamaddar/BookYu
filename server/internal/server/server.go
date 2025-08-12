package server

import (
	"fmt"

	"github.com/RonitSamaddar/BookYu/internal/config"
)

func Start() error {
	if err := config.GetConfigInputs(); err != nil {
		return nil
	}

	if err := setup(); err != nil {
		return err
	}

	fmt.Println("Server started")
	return nil
}

func setup() error {
	fmt.Println("Server setup done")
	return nil
}
