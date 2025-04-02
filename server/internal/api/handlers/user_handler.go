package handlers

import (
	"server/internal/db/repositories"
	"server/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserRepo *repositories.UserRepository
}

func NewUserHandler(userRepo *repositories.UserRepository) *UserHandler {
	return &UserHandler{UserRepo: userRepo}
}

// GetUsers handles GET /users request
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.UserRepo.GetAllUsers()
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

	user.ID = uuid.New()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)

	if err := h.UserRepo.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user ",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created successfully",
	})
}

// GetUser handles GET /users/:id request
func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.UserRepo.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve user",
		})
	}

	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	return c.JSON(user)
}

// UpdateUser handles PUT /users/:id request
func UpdateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "User updated"})
}

// DeleteUser handles DELETE /users/:id request
func DeleteUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "User deleted"})
}
