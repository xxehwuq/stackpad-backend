package main

import (
	"log"

	"github.com/yaroslavyarosh/stackpad-backend/config"
	"github.com/yaroslavyarosh/stackpad-backend/internal/app"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal("error initializing config: ", err)
	}

	app.Run(cfg)
}
