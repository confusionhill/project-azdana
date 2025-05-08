package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Server   ServerConfig `json:"server"`
	Settings Settings     `json:"settings"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("private/aqw.config.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config JSON: %w", err)
	}

	return &cfg, nil
}

type ServerConfig struct {
	Database DatabaseConfig `json:"database"`
	Name     string         `json:"name"`
	WebPort  string         `json:"webPort"`
	GamePort string         `json:"gamePort"`
}

type DatabaseConfig struct {
	Type string `json:"type"`
	Host string `json:"host"`
}

type Settings struct {
	Game GameSetting `json:"game"`
}

type GameSetting struct {
	Client     string `json:"client"`
	Title      string `json:"title"`
	Background string `json:"background"`
}
