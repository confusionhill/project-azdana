package main

import (
	"log"

	"com.github/confusionhill-aqw-ps/application"
	webbackend "com.github/confusionhill-aqw-ps/application/webBackend"
	"com.github/confusionhill-aqw-ps/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	_, _, _, handlers, err := application.RunApplication(cfg)
	if err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
	err = webbackend.RunWebBackendApp(cfg, handlers)
	if err != nil {
		log.Fatalf("Failed to run web backend: %v", err)
	}
}
