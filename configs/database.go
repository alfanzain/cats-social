package configs

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func DatabaseConnect() error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err = sql.Open("postgres", url)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}
