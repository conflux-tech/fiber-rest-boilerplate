package providers

import "github.com/conflux-tech/fiber-rest-boilerplate/config"

var appConfig *config.Configuration

// SetConfig func
func SetConfig(c *config.Configuration) {
	appConfig = c
}

// GetConfig func
func GetConfig() (c *config.Configuration) {
	return appConfig
}
