package repository

import "golang-grpc-sqlboiler-mysql/pkg/models"

type IUserRepositories interface {
	GetUser(userID int) (*models.MUser, error)
}
