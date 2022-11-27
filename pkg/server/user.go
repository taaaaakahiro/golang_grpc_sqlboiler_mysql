package server

import (
	"context"
	grpcpb "golang-grpc-sqlboiler-mysql/pkg/grpc"
	"strconv"

	"github.com/friendsofgo/errors"
)

func (s *Server) User(ctx context.Context, req *grpcpb.UserRequest) (*grpcpb.UserResponse, error) {
	id, _ := strconv.Atoi(req.Id)
	user, err := s.repo.User.GetUser(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &grpcpb.UserResponse{
		Id:   int32(user.ID),
		Name: user.Name,
		Age:  int32(user.Age),
	}, nil
}
