package service

import (
	"github.com/alexflint/go-arg"
)

type Config struct {
	GrpcPort     string `arg:"env:GRPC_PORT" default:":50051"`
	ElasticURL   string `arg:"env:ELASTIC_URL,required"`
	ElasticIndex string `arg:"env:ELASTIC_INDEX" default:"explore"`
}

func NewConfig() (Config, error) {
	config := Config{}
	arg.MustParse(&config)

	return config, nil
}
