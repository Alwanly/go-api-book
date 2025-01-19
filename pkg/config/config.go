package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server struct {
		Host string `envconfig:"SERVER_HOST" default:"localhost"`
		Port int    `envconfig:"SERVER_PORT"`
	}

	Authentication struct {
		BasicUsername string `envconfig:"AUTH_BASIC_USERNAME"`
		BasicPassword string `envconfig:"AUTH_BASIC_PASSWORD"`
		JwtSecret     string `envconfig:"JWT_SECRET"`
	}

	Database struct {
		URI string `envconfig:"DB_URI"`
	}

	Redis struct {
		URI string `envconfig:"REDIS_URI"`
	}

	Apm struct {
		Active bool `envconfig:"ELASTIC_APM_ACTIVE" default:"false"`
	}
}

var GlobalConfig Config = Config{}

func LoadConfig() {
	godotenv.Load()
	envconfig.Process("", &GlobalConfig)
}
