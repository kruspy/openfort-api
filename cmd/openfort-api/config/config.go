package config

import (
	"encoding/json"
	"fmt"
	"openfort-api/cmd/openfort-api/logger"
	"os"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

// Config holds the application configuration data.
type Config struct {
	// Server
	Address string `json:"address"  envconfig:"ADDRESS"`
	Port    string `json:"port"     envconfig:"PORT"`

	// Database
	DbUrl string `json:"db_url" envconfig:"DB_URL"`
}

var cfg *Config

// GetConfig returns the current app config.
func GetConfig() *Config {
	return cfg
}

// Load reads config from the config file and parses it into memory.
func Load() {
	cfg = &Config{}

	path := os.Getenv("CONFIG")
	if path != "" {
		// If a config file is set, use it.
		if err := loadFromFile(path, cfg); err != nil {
			logger.Fatal(
				fmt.Sprintf("error loading config from file: %s", err.Error()),
				zap.String("path", path),
			)
		}

	} else {
		// Default to env. variables.
		if err := loadFromEnv(cfg); err != nil {
			logger.Fatal(fmt.Sprintf("error loading config from env: %s", err.Error()))
		}
	}
}

// loadFromFile reads the config from a file.
func loadFromFile(path string, config *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, config)
}

// loadFromEnv reads the config from env variables.
func loadFromEnv(config *Config) error {
	return envconfig.Process("", config)
}
