package service

import (
	"log"

	"google.golang.org/grpc"

	"muzz/elastic"
	"muzz/explore"
)

func (s *Service) setupGrpc(logger *log.Logger, esClient *elastic.Client) error {
	g := grpc.NewServer()

	grpcService, err := explore.NewService(logger, esClient)
	if err != nil {
		return err
	}

	explore.RegisterExploreServiceServer(g, grpcService)

	s.grpcServer = g

	return nil
}
