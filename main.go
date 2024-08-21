package main

import (
	"hcs-ocs/config"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting ocs process")
	log.Info("Loading configuration")
	err := config.LoadConfig();
	if err != nil {
		log.Fatalf("There was an error loading config: %s", err.Error())
	}
	log.Info("Configuration parsed successfully")
}
