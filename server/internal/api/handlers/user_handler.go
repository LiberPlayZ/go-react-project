package handlers

import (
	"server/internal/api/services"
	"server/internal/models"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// GetUsers handles GET /users request
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.UserService.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}
	return c.JSON(users)
}

// CreateUser handles POST /users request
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// validate not null fields

	if user.Username == "" || user.Password == "" || user.Role == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing request fields",
		})
	}

	if err := h.UserService.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user ",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created successfully",
	})
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
