package server

import (
	"context"
	"database/sql"
	grpcpb "golang-grpc-sqlboiler-mysql/pkg/grpc"
	"time"

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

func (s *Server) UserServerStream(req *grpcpb.UserRequest, stream grpcpb.UserService_UserServerStreamServer) error {
	resCount := 5
	for i := 1; i < resCount; i++ {
		user, err := s.repo.User.GetUser(i)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				return nil
			default:
				return errors.WithStack(err)
			}
		}
		if err := stream.Send(&grpcpb.UserResponse{
			Id:   int32(user.ID),
			Name: user.Name,
			Age:  int32(user.Age),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}
