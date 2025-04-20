package routes

import (
	"database/sql"
	"server/internal/api/handlers"
	"server/internal/auth"
	"server/internal/db/repositories"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes registers all API routes
func SetupRoutes(app *fiber.App, db *sql.DB, hashingCost int) {
	userRepo := repositories.NewUserRepository(db)
	hashing := auth.NewHashing(hashingCost)
	userHandler := handlers.NewUserHandler(userRepo, hashing)

	todoRepo := repositories.NewTodoRepository(db)
	todoHandler := handlers.NewTodoHandler(todoRepo)

	api := app.Group("/api") // API versioning
	// users gateway

	api.Get("/users", userHandler.GetUsers)
	api.Post("/users", userHandler.CreateUser)
	api.Get("/users/:id", userHandler.GetUserById) // <-- New route
	api.Post("/users/login", userHandler.Login)
	// api.Put("/users/:id", handlers.UpdateUser)
	// api.Delete("/users/:id", handlers.DeleteUser)

	// todos gateway

	api.Get("/todos", todoHandler.GetTodos)
	api.Post("/todos", todoHandler.CreateTodo)
	api.Put("/todos/:id", todoHandler.UpdateTodoToCompleted)
	api.Delete("/todos/:id", todoHandler.DeleteTodo)

}
