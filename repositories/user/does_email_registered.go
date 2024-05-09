package userrepo

import (
	"catssocial/configs"
	"database/sql"
	"errors"
	"log"
)

func (u *UserRepository) DoesEmailRegistered(email string) (bool, error) {
	var scannedEmail string
	err := configs.DB.QueryRow(`SELECT email FROM users WHERE email = $1`, email).Scan(&scannedEmail)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Fatalln(err)
		return false, err
	}

	if len(scannedEmail) == 0 {
		return false, nil 
	}

	return true, nil
}
