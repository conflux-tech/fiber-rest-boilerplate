package config

import (
	"github.com/spf13/viper"
)

// DBConfig is the database configuration definition
type DBConfig struct {
	Enabled		bool
	Driver		string
	Host		string
	Port		int
	User		string
	Password	string
	Database	string
	SSLMode		string
}

func loadDBConfig() (config DBConfig, err error) {
	provider := viper.New()
	provider.SetEnvPrefix("DB")
	provider.AutomaticEnv()
	setDefaultDBConfig(provider)

	err = provider.Unmarshal(&config)

	return config, err
}

func setDefaultDBConfig(provider *viper.Viper) {
	provider.SetDefault("Enabled", true)
	provider.SetDefault("Driver", "pg")
	provider.SetDefault("Host", "localhost")
	provider.SetDefault("Port", 5432)
	provider.SetDefault("User", "postgres")
	provider.SetDefault("Database", "postgres")
	provider.SetDefault("Password", "")
	provider.SetDefault("SSLMode", "disable")
}
