package main

import (
	"log"

	"github.com/anazibinurasheed/project-device-mart/pkg/config"
	"github.com/anazibinurasheed/project-device-mart/pkg/di"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	server.Start(config.PORT)
}
