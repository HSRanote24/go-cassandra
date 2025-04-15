package routes

import (
	"go-cassandra/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {
	user := app.Group("/users")
	user.Post("/", handlers.CreateUser)
	user.Get("/", handlers.GetUsers)
	user.Get("/:id", handlers.GetUser)
	user.Put("/:id", handlers.UpdateUser)
	user.Delete("/:id", handlers.DeleteUser)
}
