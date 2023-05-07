package main

import "github.com/caarlos0/env/v6"

type Config struct {
	LogLevel      string `env:"WEATHER_CLIENT_LOG_LEVEL" envDefault:"debug"`
	Host          string `env:"WEATHER_CLIENT_HOST" envDefault:"localhost"`
	Port          string `env:"WEATHER_CLIENT_PORT" envDefault:"9090"`
	ServerAddress string `env:"WEATHER_SERVER_ADDRESS,required"`
}

func New() (Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
