package main

import (
	"github.com/felipematheus1337/Gopportunities/config"
	"github.com/felipematheus1337/Gopportunities/router"
)

var (
	logger config.Logger
)

func main() {

	logger = *config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error", err)
		return
	}

	router.Initialize()
}
