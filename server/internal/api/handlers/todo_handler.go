package handlers

import (
	"log"
	"server/internal/db/repositories"
	"server/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// validate not null fields

	if todo.Title == "" || todo.Description == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing request fields",
		})
	}

	todo.ID = uuid.New()

	createdTodo, err := h.TodoRepo.CreateTodo(todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create todo ",
		})
	}

	return c.JSON(createdTodo)

}
