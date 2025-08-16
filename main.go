package main

import (
	"github.com/DaviRodrigues/Project-go-with-gin/config"
	"github.com/DaviRodrigues/Project-go-with-gin/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Intialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		panic(err)
	}

	router.Initialize()
}
