package main

import (
	"cats-social/routes"
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "local_cats_social"
)

func Connect() error {
	// var err error
	// db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	// if err != nil {
	// 	return err
	// }
	// if err = db.Ping(); err != nil {
	// 	return err
	// }
	return nil
}

func main() {
	// Connect with database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()

	// Register routes
	routes.RouteRegister(app)

	// Start the server on port 3000
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
