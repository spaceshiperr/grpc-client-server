package configs

import "github.com/caarlos0/env/v6"

type Config struct {
	LogLevel string `env:"WEATHER_SERVER_LOGLEVEL" envDefault:"debug"`
	Host     string `env:"WEATHER_SERVER_HOST" envDefault:"localhost"`
	Port     string `env:"WEATHER_SERVER_PORT" envDefault:"8080"`

	OwmApiKey string `env:"OWM_API_KEY,required"`
}

func New() (Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
