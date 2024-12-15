package service

import (
	"context"
	"log"
	"muzz/elastic"
	"net"
	"os"

	esgo "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"google.golang.org/grpc"
)

// Service definition
type Service struct {
	logger      *log.Logger
	grpcServer  *grpc.Server
	stopChannel chan bool
}

// NewService creates a new service
func NewService(_ context.Context, config Config) (*Service, error) {
	logger := log.New(os.Stdout, "muzz-task:", log.LstdFlags)

	esAPI, err := createElasticAPI(config.ElasticURL)
	if err != nil {
		return nil, err
	}
	esClient := elastic.NewClient(logger, esAPI)

	s := Service{
		logger:      logger,
		stopChannel: make(chan bool),
	}
	s.setupGrpc(logger, esClient)

	return &s, nil
}

// Start starts the service
func (s *Service) Start() error {
	var err error
	s.logger.Println("starting")
	// start grpc server and listener
	go func() {
		s.logger.Println("Starting gRPC server")
		l, err := net.Listen("tcp", ":50051")
		if err != nil {
			s.stopChannel <- true // stop the service
		}
		err = s.grpcServer.Serve(l)
		if err != nil {
			s.stopChannel <- true
		}
	}()
	return err
}

// Wait for the service
func (s *Service) Wait() {
	s.logger.Println("waiting")
	<-s.stopChannel
}

// Close closes the service
func (s *Service) Close() {
	s.logger.Println("closing")
	s.grpcServer.Stop()
	close(s.stopChannel)
}

func createElasticAPI(elasticURL string) (*esapi.API, error) {
	esClient, err := esgo.NewClient(esgo.Config{
		Addresses: []string{elasticURL},
	})
	if err != nil {
		return nil, err
	}
	return esClient.API, nil
}
