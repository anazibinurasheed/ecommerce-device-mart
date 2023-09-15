package main

import (
	"log"

	"github.com/anazibinurasheed/project-device-mart/pkg/config"
	"github.com/anazibinurasheed/project-device-mart/pkg/db"
	"github.com/anazibinurasheed/project-device-mart/pkg/di"
	"github.com/anazibinurasheed/project-device-mart/pkg/util/helper"
)

func main() {

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		if err := helper.SetupDB(db.GetDBInstance()); err != nil {
			log.Fatal(err)

		}

	}

	server.Start()

}
