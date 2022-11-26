package server

import (
	"context"
	"fmt"
	grpcpb "golang-grpc-sqlboiler-mysql/pkg/grpc"
	"strconv"

	"github.com/friendsofgo/errors"
)

func (s *Server) Hello(ctx context.Context, req *grpcpb.HelloRequest) (*grpcpb.HelloResponse, error) {
	id, _ := strconv.Atoi(req.Name)
	user, err := s.repo.User.GetUser(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &grpcpb.HelloResponse{
		Message: fmt.Sprintf("ID:%d, Name:%s, Age:%d", user.ID, user.Name, user.Age),
	}, nil
}
