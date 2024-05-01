package main

import (
	"database/sql"
	"fmt"
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

	// Routes

	/// Authentication & Authorization
	userGroupRoutes := app.Group("/v1/user")

	// POST /v1/user/register
	userGroupRoutes.Post("/register", func(c *fiber.Ctx) error {
		return c.Status(201).JSON("POST /v1/user/register")
	})

	// POST /v1/user/login
	userGroupRoutes.Post("/login", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("POST /v1/user/login")
	})

	/// Manage Cats
	catGroupRoutes := app.Group("/v1/cat")

	// POST /v1/cat
	catGroupRoutes.Post("", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("POST /v1/cat")
	})

	// GET /v1/cat
	catGroupRoutes.Get("", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("GET /v1/cat")
	})

	// PUT /v1/cat/{id}
	catGroupRoutes.Put("/:id", func(c *fiber.Ctx) error {
		response := fmt.Sprintf("PUT /v1/cat/%s", c.Params("id"))
		return c.Status(200).JSON(response)
	})

	// DELETE /v1/cat/{id}
	catGroupRoutes.Delete("/:id", func(c *fiber.Ctx) error {
		response := fmt.Sprintf("DELETE /v1/cat/%s", c.Params("id"))
		return c.Status(200).JSON(response)
	})

	// Start the server on port 3000
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
