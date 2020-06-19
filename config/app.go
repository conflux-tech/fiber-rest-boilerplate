package config

import (
	"github.com/spf13/viper"
)

// AppConfig is the application's configuration definition
type AppConfig struct {
	Env			string
	Port		int
}

func loadAppConfig() (config AppConfig, err error) {
	provider := viper.New()
	provider.SetEnvPrefix("APP")
	provider.AutomaticEnv()

	setDefaultAppConfig(provider)

	err = provider.Unmarshal(&config)

	return config, err
}

func setDefaultAppConfig(provider *viper.Viper) {
	provider.SetDefault("Env", "production")
	provider.SetDefault("Port", 3000)
}
