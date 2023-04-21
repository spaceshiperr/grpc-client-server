package main

import (
	"github.com/rs/zerolog/log"

	"github.com/spaceshiperr/grpc-client-server/server/internal/app"
	"github.com/spaceshiperr/grpc-client-server/server/internal/configs"
)

func main() {
	cfg, err := configs.New()
	if err != nil {
		log.Fatal().Msgf("Configuration failed: %v", err.Error())
	}

	app.Run(cfg)
}
