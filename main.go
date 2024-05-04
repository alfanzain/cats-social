package main

import (
	"catssocial/configs"
	"catssocial/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	if err := configs.DatabaseConnect(); err != nil {
		log.Fatal("Cannot connect database: ", err)
	}

	app := fiber.New()
	routes.SetupRoutes(app)

	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	err = app.Listen(":" + config.APPPort)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
