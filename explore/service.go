package explore

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service provides a wrapper for the Explore grpc server
type Service struct {
	UnimplementedExploreServiceServer
}

func NewService() (*Service, error) {
	return &Service{}, nil
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
