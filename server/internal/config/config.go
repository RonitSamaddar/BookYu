package config

import (
	"log"
	"os"

	"github.com/RonitSamaddar/BookYu/internal/consts"
	"github.com/RonitSamaddar/BookYu/internal/types"
)

func LoadInputConfigs() (*types.Config, error) {
	config := types.Config{}
	loadEnvs(&config)
	log.Println("Loading input configs")
	return &config, nil
}

func loadEnvs(config *types.Config) {
	config.GoogleBooksApiKey = os.Getenv(consts.GoogleBooksApiKeyEnvVar)
}
