package main

import (
	"log"

	webbackend "com.github/confusionhill-aqw-ps/application/webBackend"
	"com.github/confusionhill-aqw-ps/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	webbackend.RunWebBackendApp(cfg)
}
