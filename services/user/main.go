package userservice

import (
	userrepo "catssocial/repositories/user"
	"errors"
)

type UserService struct {
	UserRepository userrepo.IUserRepository
}

type IUserService interface {
	Register(*RegisterPayload) (*RegisterResult, error)
}

var ErrEmailAlreadyRegistered = errors.New("email sudah digunakan")

func New(userRepo userrepo.IUserRepository) IUserService {
	return &UserService{
		UserRepository: userRepo,
	}
}
