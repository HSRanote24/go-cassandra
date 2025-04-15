package main

import (
	"go-cassandra/config"
	"go-cassandra/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitCassandra()
	defer config.Session.Close()

	app := fiber.New()
	routes.RegisterUserRoutes(app)

	app.Listen(":3000")
}
