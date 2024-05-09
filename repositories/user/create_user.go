package userrepo

import (
	"catssocial/configs"
	"catssocial/entities"
	"log"
)

type (
	RegisterPayload struct {
		Email    string
		Name     string
		Password string
	}
)

func (u *UserRepository) CreateUser(p *RegisterPayload) (*entities.User, error) {
	query := `INSERT INTO users (email, name, password) VALUES ($1, $2, $3) RETURNING id`

	statement, err := configs.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	var ID string
	err = statement.QueryRow(p.Email, p.Name, p.Password).Scan(&ID)
	if err != nil {
		log.Printf("Error inserting user: %s", err)
		return nil, err
	}

	return &entities.User{
		ID:    ID,
		Name:  p.Name,
		Email: p.Email,
	}, nil
}
