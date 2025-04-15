package handlers

import (
	"go-cassandra/models"
	"go-cassandra/services"

	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	u := new(models.User)
	if err := c.BodyParser(u); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := services.CreateUser(u); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}
	return c.Status(201).JSON(u)
}

func GetUsers(c *fiber.Ctx) error {
	users, err := services.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error retrieving users"})
	}
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id, err := gocql.ParseUUID(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"})
	}
	user, err := services.GetUserByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := gocql.ParseUUID(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"})
	}
	u := new(models.User)
	if err := c.BodyParser(u); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := services.UpdateUser(id, u); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Update failed"})
	}
	u.ID = id
	return c.JSON(u)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := gocql.ParseUUID(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"})
	}
	if err := services.DeleteUser(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Delete failed"})
	}
	return c.SendStatus(204)
}
