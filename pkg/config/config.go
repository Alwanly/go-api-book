package config

import (
	"github.com/spf13/viper"
)

func LoadConfig(configName string) (*GlobalConfig, error) {
	var cfg GlobalConfig

	// Load config from file
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(configName)

	// load from environment variables
	viper.AutomaticEnv()

	return &cfg, nil
}
