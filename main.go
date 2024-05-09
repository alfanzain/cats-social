package main

import (
	"catssocial/configs"
	userhandler "catssocial/handlers/user"
	"catssocial/middlewares"
	"catssocial/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	if err := configs.DatabaseConnect(); err != nil {
		log.Fatal("Cannot connect database: ", err)
	}

	app := fiber.New()

	routeConfig := routes.RouteConfig{
		App:            app,
		UserHandler:    userhandler.New(config.Logger),
		AuthMiddleware: middlewares.Authenticated(),
	}
	routeConfig.Setup()

	err = app.Listen(":" + config.APPPort)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
