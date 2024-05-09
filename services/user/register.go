package userservice

import (
	"catssocial/helpers"
	userrepo "catssocial/repositories/user"
)

type (
	RegisterPayload struct {
		Email    string
		Name     string
		Password string
	}

	RegisterResult struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		AccessToken string `json:"accessToken"`
	}
)

func (u *UserService) Register(p *RegisterPayload) (*RegisterResult, error) {
	emailRegistered, err := u.UserRepository.DoesEmailRegistered(p.Email)
	if err != nil {
		return nil, err
	}
	if emailRegistered {
		return nil, ErrEmailAlreadyRegistered
	}

	p.Password, _ = helpers.HashPassword(p.Password)

	user, err := u.UserRepository.CreateUser(&userrepo.RegisterPayload{
		Email:    p.Email,
		Name:     p.Name,
		Password: p.Password,
	})

	accessToken, err := helpers.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &RegisterResult{
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: accessToken,
	}, nil
}
