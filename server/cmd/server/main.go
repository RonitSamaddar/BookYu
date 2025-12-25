package main

import (
	"github.com/RonitSamaddar/BookYu/internal/server"
)

func main() {
	if err:= server.Start(); err != nil {
		panic(err)
	}
}
