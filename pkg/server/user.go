package server

import (
	"context"
	"database/sql"
	grpcpb "golang-grpc-sqlboiler-mysql/pkg/grpc"

	"github.com/friendsofgo/errors"
)

func (s *Server) User(ctx context.Context, req *grpcpb.UserRequest) (*grpcpb.UserResponse, error) {
	user, err := s.repo.User.GetUser(int(req.Id))
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, errors.WithStack(err)
		}
	}

	return &grpcpb.UserResponse{
		Id:   int32(user.ID),
		Name: user.Name,
		Age:  int32(user.Age),
	}, nil
}
