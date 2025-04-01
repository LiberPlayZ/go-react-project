package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// GetUsers handles GET /users request
func GetUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get all users"})
}

// CreateUser handles POST /users request
func CreateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "User created"})
}

// GetUser handles GET /users/:id request
func GetUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Get single user"})
}

// UpdateUser handles PUT /users/:id request
func UpdateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "User updated"})
}

// DeleteUser handles DELETE /users/:id request
func DeleteUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "User deleted"})
}
