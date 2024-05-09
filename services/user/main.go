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
	Login(*LoginPayload) (*LoginResult, error)
}

var (
	ErrEmailAlreadyRegistered = errors.New("email sudah digunakan")
	ErrUserNotFound           = errors.New("user tidak ditemukan")
	ErrInvalidPassword        = errors.New("password salah")
)

func New(userRepo userrepo.IUserRepository) IUserService {
	return &UserService{
		UserRepository: userRepo,
	}
}
