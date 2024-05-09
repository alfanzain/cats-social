package userrepo

import (
	"catssocial/configs"
	"catssocial/entities"
	"database/sql"
	"errors"
	"log"
)

func (u *UserRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := configs.DB.QueryRow(`SELECT id, email, password FROM users WHERE email = $1`, email).Scan(&user.ID, &user.Email, &user.Password)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Fatalln(err)
		return nil, err
	}

	return &user, nil
}
