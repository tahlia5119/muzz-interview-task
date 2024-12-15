package service

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

// Service definition
type Service struct {
	logger      *log.Logger
	grpcServer  *grpc.Server
	stopChannel chan bool
}

// NewService creates a new service
func NewService(_ context.Context) *Service {
	logger := log.New(os.Stdout, "muzz-task:", log.LstdFlags)
	s := Service{
		logger:      logger,
		stopChannel: make(chan bool),
	}
	s.setupGrpc()
	return &s
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
