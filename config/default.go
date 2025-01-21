package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

func loadDefaults() {
	// global config
	viper.SetDefault("ENV", "development")
	viper.SetDefault("PORT", 9000)
	viper.SetDefault("PORT_GRPC", 9001)
	viper.SetDefault("LOG_LEVEL", "debug")

	// set service name and version
	viper.SetDefault("SERVICE_NAME", "go-codebase")
	if versionBytes, err := os.ReadFile("VERSION"); err == nil {
		// remove the newline character
		viper.SetDefault("SERVICE_VERSION", strings.TrimSuffix(string(versionBytes), "\n"))
	} else {
		viper.SetDefault("SERVICE_VERSION", "0.0.1")
	}

	// postgres default
	viper.SetDefault("POSTGRES_URI", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	viper.SetDefault("POSTGRES_MAX_OPEN_CONNECTIONS", 10)
	viper.SetDefault("POSTGRES_MAX_IDLE_CONNECTIONS", 5)

	// redis default
	viper.SetDefault("REDIS_URI", "redis://redis:6379/0")
}
