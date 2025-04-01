package routes

import (
	"database/sql"
	"server/internal/api/handlers"
	"server/internal/api/services"
	"server/internal/db/repositories"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes registers all API routes
func SetupRoutes(app *fiber.App, db *sql.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	api := app.Group("/api") // API versioning

	api.Get("/users", userHandler.GetUsers)
	api.Post("/users", userHandler.CreateUser)
	// api.Post("/users", handlers.CreateUser)
	// api.Get("/users/:id", handlers.GetUser)
	// api.Put("/users/:id", handlers.UpdateUser)
	// api.Delete("/users/:id", handlers.DeleteUser)
}
