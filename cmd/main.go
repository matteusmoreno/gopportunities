package main

import (
	"github.com/matteusmoreno/gopportunities/config"
	"github.com/matteusmoreno/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger()

	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	router.Initialize()
}
