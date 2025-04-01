package routes

import (
	"server/internal/api/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes registers all API routes
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api") // API versioning

	api.Get("/users", handlers.GetUsers)
	api.Post("/users", handlers.CreateUser)
	api.Get("/users/:id", handlers.GetUser)
	api.Put("/users/:id", handlers.UpdateUser)
	api.Delete("/users/:id", handlers.DeleteUser)
}
