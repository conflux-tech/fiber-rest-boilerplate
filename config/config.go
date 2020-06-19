package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Configuration is the project configuration definition
type Configuration struct {
	Application	AppConfig
	Database	DBConfig
}

// LoadConfigs loads configs/envvars
func LoadConfigs() (config Configuration, err error) {
	if os.Getenv("APP_ENV") == "development" {
		dotenvErr := godotenv.Load()
		if dotenvErr != nil {
			panic(dotenvErr)
		}
	}

	// Load Application Configs
	appConfig, _ := loadAppConfig()
	config.Application = appConfig

	// Load Database Configs
	dbConfig, _ := loadDBConfig()
	config.Database = dbConfig

	return config, nil
}
