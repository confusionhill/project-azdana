package main

import (
	"log"

	"com.github/confusionhill-aqw-ps/application"
	"com.github/confusionhill-aqw-ps/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	err = application.RunApplication(cfg)
	if err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}
