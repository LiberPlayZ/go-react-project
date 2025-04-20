package handlers

import (
	"server/internal/auth"
	"server/internal/db/repositories"
	"server/internal/dtos"
	"server/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	UserRepo       *repositories.UserRepository
	HashingHandler *auth.Hashing
}

func NewUserHandler(userRepo *repositories.UserRepository, hashing *auth.Hashing) *UserHandler {
	return &UserHandler{UserRepo: userRepo, HashingHandler: hashing}
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

	if user.Username == "" || user.Email == "" || user.Password == "" || user.Role == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing request fields",
		})
	}

	user.ID = uuid.New()

	hashPassword, err := h.HashingHandler.HashPassword(user.Password)
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

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var userLogin dtos.UserLoginDto

	// check if request has body
	if err := c.BodyParser(&userLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	//check if email and password are not null .
	if userLogin.Email == "" || userLogin.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing request fields",
		})
	}

	user, err := h.UserRepo.GetUserByEmail(userLogin.Email)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Failed to retrieve user",
		})

	}

	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	match := h.HashingHandler.VerifyPassword(userLogin.Password, user.Password)
	if !match {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "password is invalid",
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
