package userservice

import (
	"catssocial/helpers"
)

type (
	LoginPayload struct {
		Email    string
		Name     string
		Password string
	}

	LoginResult struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		AccessToken string `json:"accessToken"`
	}
)

func (u *UserService) Login(p *LoginPayload) (*LoginResult, error) {

	user, err := u.UserRepository.FindByEmail(p.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}

	isValidPassword := helpers.CheckPasswordHash(p.Password, user.Password)
	if !isValidPassword {
		return nil, ErrInvalidPassword
	}

	accessToken, err := helpers.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &LoginResult{
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: accessToken,
	}, nil
}
