package main

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/spaceshiperr/grpc-client-server/proto/generated"
)

type WeatherServiceClient struct {
	Client pb.WeatherServiceClient
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cfg, err := New()
	if err != nil {
		log.Fatal().Msgf("configuration failed: %v", err.Error())
	}

	level, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatal().Msgf("failed to parse config, error: %s", err.Error())
	}

	logger := log.Logger.WithLevel(level)

	conn, err := grpc.Dial(cfg.ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Msgf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := WeatherServiceClient{
		Client: pb.NewWeatherServiceClient(conn),
	}

	if err := client.GetWeatherByCity(ctx); err != nil {
		logger.Msgf("err: %s", err.Error())
	}
}

func (c WeatherServiceClient) GetWeatherByCity(ctx context.Context) error {
	in := &pb.GetWeatherByCityRequest{
		City: "St. Petersburg",
	}

	out, err := c.Client.GetWeatherByCity(ctx, in)
	if err != nil {
		return err
	}

	fmt.Printf("Temp: %v, TempMin: %v, TempMax: %v, FeelsLike: %v\n", out.Temp, out.TempMin, out.TempMax, out.FeelsLike)
	fmt.Printf("Pressure: %v, Humidity: %v, WindSpeed: %v\n", out.Pressure, out.Humidity, out.WindSpeed)
	fmt.Printf("Conditions: %s, Description: %v\n", out.Conditions, out.Description)

	return nil
}
