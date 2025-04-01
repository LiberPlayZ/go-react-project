package server

import (
	"fmt"
	"log"
	"server/config"

	"github.com/gofiber/fiber/v2"
)

// NewServer initializes and returns a new Fiber app
func NewServer(port string) *fiber.App {
	_, err := config.GetIntEnv(port)
	if err != nil {
		log.Print("Failed to load port from config . using deffualt port 4000")
		port = "4000"
	}
	app := fiber.New()

	// // Register routes
	// routes.SetupRoutes(app)

	// Start server
	app.Listen(":" + port)
	fmt.Printf("🚀 Server is running on port %s\n", port)
	return app
}
