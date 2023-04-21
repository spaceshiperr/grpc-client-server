package configs

import "github.com/caarlos0/env/v6"

type Config struct {
	LogLevel string `env:"WEATHER_LOGLEVEL" envDefault:"debug"`
	Host     string `env:"WEATHER_HOST" envDefault:"localhost"`
	Port     string `env:"WEATHER_PORT" envDefault:"80"`

	OwmApiKey string `env:"OWM_API_KEY,required"`
}

func New() (*Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
