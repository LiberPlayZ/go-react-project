package routes

import (
	"database/sql"
	"server/internal/api/handlers"
	"server/internal/db/repositories"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes registers all API routes
func SetupRoutes(app *fiber.App, db *sql.DB) {
	userRepo := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	todoRepo := repositories.NewTodoRepository(db)
	todoHandler := handlers.NewTodoHandler(todoRepo)

	api := app.Group("/api") // API versioning
	// users gateway

	api.Get("/users", userHandler.GetUsers)
	api.Post("/users", userHandler.CreateUser)
	api.Get("/users/:id", userHandler.GetUserById) // <-- New route
	// api.Put("/users/:id", handlers.UpdateUser)
	// api.Delete("/users/:id", handlers.DeleteUser)

	// todos gateway

	api.Get("/todos", todoHandler.GetTodos)
	api.Post("/todos", todoHandler.CreateTodo)

}
