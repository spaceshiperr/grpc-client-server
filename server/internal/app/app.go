package app

import (
	"github.com/rs/zerolog"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	pb "github.com/spaceshiperr/grpc-client-server/proto/generated"
	"github.com/spaceshiperr/grpc-client-server/server/internal/configs"
	weatherGrpc "github.com/spaceshiperr/grpc-client-server/server/internal/controller/grpc"
)

func Run(cfg configs.Config) {
	level, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatal().Msgf("failed to parse config, error: %s", err.Error())
	}

	logger := log.Logger.WithLevel(level)

	address := net.JoinHostPort(cfg.Host, cfg.Port)

	listen, err := net.Listen("tcp", address)
	if err != nil {
		logger.Msgf("failed to listen to %v, error: %v", address, err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterWeatherServiceServer(s, weatherGrpc.NewWeatherServer(cfg, logger))

	logger.Msgf("start GRPC on %s port", cfg.Port)

	if err := s.Serve(listen); err != nil {
		logger.Msgf("failed to serve: %v", err.Error())
	}
}
