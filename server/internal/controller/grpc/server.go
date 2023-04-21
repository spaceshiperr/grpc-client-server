package grpc

import (
	pb "github.com/spaceshiperr/grpc-client-server/proto"
)

type WeatherServer struct {
	pb.UnimplementedWeatherServe
}
