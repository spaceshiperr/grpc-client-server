package grpc

import (
	"context"

	owm "github.com/briandowns/openweathermap"
	"github.com/rs/zerolog"

	pb "github.com/spaceshiperr/grpc-client-server/proto/generated"
	"github.com/spaceshiperr/grpc-client-server/server/internal/configs"
)

type WeatherServer struct {
	pb.UnimplementedWeatherServiceServer
	cfg    configs.Config
	logger *zerolog.Event
}

func NewWeatherServer(cfg configs.Config, logger *zerolog.Event) WeatherServer {
	return WeatherServer{
		cfg:    cfg,
		logger: logger,
	}
}

func (s WeatherServer) GetWeatherByCity(ctx context.Context, in *pb.GetWeatherByCityRequest) (*pb.GetWeatherByCityResponse, error) {
	city := in.GetCity()

	w, err := owm.NewCurrent("C", "ru", s.cfg.OwmApiKey)
	if err != nil {
		return nil, err
	}

	if err = w.CurrentByName(city); err != nil {
		return nil, err
	}

	out := &pb.GetWeatherByCityResponse{
		Temp:      float32(w.Main.Temp),
		FeelsLike: float32(w.Main.FeelsLike),
		Pressure:  uint32(w.Main.Pressure),
		Humidity:  uint32(w.Main.Humidity),
		TempMin:   float32(w.Main.TempMin),
		TempMax:   float32(w.Main.TempMax),
		WindSpeed: uint32(w.Wind.Speed),
	}

	if len(w.Weather) > 0 {
		out.Conditions = w.Weather[0].Main
		out.Description = w.Weather[0].Description
	}

	return out, nil
}
