package server

import (
	grpcpb "golang-grpc-sqlboiler-mysql/pkg/grpc"
	"golang-grpc-sqlboiler-mysql/pkg/infrastracture/persistence"
)

type Server struct {
	repo *persistence.Repositories
	grpcpb.UnimplementedGreetingServiceServer
}

func NewServer(r *persistence.Repositories) (*Server, error) {
	return &Server{
		repo: r,
	}, nil
}
