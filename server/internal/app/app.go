package app

import (
	"github.com/rs/zerolog/log"
	"github.com/spaceshiperr/grpc-client-server/server/internal/configs"
	"google.golang.org/grpc"
	"net"
)

func Run(cfg *configs.Config) {
	logger := log.Logger

	address := net.JoinHostPort(cfg.Host, cfg.Port)

	listen, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatal().Msgf("Failed to listen to %v, error: %v", address, err.Error())
	}

	s := grpc.NewServer()
	s.RegisterService()

}
