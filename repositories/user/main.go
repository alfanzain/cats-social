package userrepo

import "catssocial/entities"

type UserRepository struct{}

type IUserRepository interface {
	DoesEmailRegistered(string) (bool, error)
	CreateUser(*RegisterPayload) (*entities.User, error)
	FindByEmail(string) (*entities.User, error)
}

func New() IUserRepository {
	return &UserRepository{}
}
