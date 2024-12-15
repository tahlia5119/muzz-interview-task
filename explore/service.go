package explore

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"muzz/elastic"
)

// Service provides a wrapper for the Explore grpc server
type Service struct {
	UnimplementedExploreServiceServer
	logger        *log.Logger
	elasticClient *elastic.ElasticClient
}

func NewService(logger *log.Logger, elasticClient *elastic.ElasticClient) (*Service, error) {
	return &Service{
		logger:        logger,
		elasticClient: elasticClient,
	}, nil
}

func (s *Service) ListLikedYou(ctx context.Context, in *ListLikedYouRequest) (*ListLikedYouResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Method not implemented")
}

func (s *Service) ListNewLikedYou(ctx context.Context, in *ListLikedYouRequest) (*ListLikedYouResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Method not implemented")
}

func (s *Service) CountLikedYou(ctx context.Context, in *CountLikedYouRequest) (*CountLikedYouResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Method not implemented")
}

func (s *Service) PutDecision(ctx context.Context, in *PutDecisionRequest) (*PutDecisionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Method not implemented")
}
