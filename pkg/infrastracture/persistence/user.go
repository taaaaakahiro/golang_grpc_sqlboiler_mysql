package persistence

import (
	"context"
	"database/sql"
	"golang-grpc-sqlboiler-mysql/pkg/domain/repository"
	"golang-grpc-sqlboiler-mysql/pkg/models"

	"github.com/friendsofgo/errors"
)

type UserRepository struct {
	ctx context.Context
	db  *sql.DB
}

var _ repository.IUserRepositories = (*UserRepository)(nil)

func NewUserRepository(ctx context.Context, db *sql.DB) *UserRepository {
	return &UserRepository{
		ctx: ctx,
		db:  db,
	}
}

func (r *UserRepository) GetUser(userID int) (*models.MUser, error) {
	user, err := models.FindMUser(r.ctx, r.db, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return user, nil

}
