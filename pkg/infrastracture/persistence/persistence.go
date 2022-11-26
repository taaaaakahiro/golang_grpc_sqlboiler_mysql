package persistence

import (
	"context"
	"database/sql"
	"golang-grpc-sqlboiler-mysql/pkg/domain/repository"
)

type Repositories struct {
	User repository.IUserRepositories
}

func NewRepositories(ctx context.Context, db *sql.DB) (*Repositories, error) {
	return &Repositories{
		User: NewUserRepository(ctx, db),
	}, nil

}
