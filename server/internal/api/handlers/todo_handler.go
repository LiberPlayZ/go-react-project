package handlers

import (
	"log"
	"server/internal/db/repositories"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	TodoRepo *repositories.TodoRepository
}

func NewTodoHandler(todoRepo *repositories.TodoRepository) *TodoHandler {
	return &TodoHandler{TodoRepo: todoRepo}
}

func (h *TodoHandler) GetTodos(c *fiber.Ctx) error {

	todos, err := h.TodoRepo.GetAllTodos()
	if err != nil {
		log.Print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve todos",
		})
	}
	return c.JSON(todos)
}
