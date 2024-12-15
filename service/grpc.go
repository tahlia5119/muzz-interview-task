package service

import (
	"google.golang.org/grpc"

	"muzz/explore"
)

func (s *Service) setupGrpc() error {
	g := grpc.NewServer()

	grpcService, err := explore.NewService()
	if err != nil {
		return err
	}

	explore.RegisterExploreServiceServer(g, grpcService)

	s.grpcServer = g

	return nil
}
