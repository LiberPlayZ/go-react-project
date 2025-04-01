package server

import (
	"fmt"
	"log"
	"server/config"
	"server/internal/api/routes"
	"server/internal/db/loaders"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// NewServer initializes and returns a new Fiber app
func NewServer() *fiber.App {

	// load config
	AppConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load config for app, error : %v", err)
	}

	// connect to db
	db, err := loaders.ConnectToDb(AppConfig)
	if err != nil {
		log.Fatalf("❌ Could not connect to database: %v", err)
	}

	defer db.Close()

	app := fiber.New()

	// Register routes
	routes.SetupRoutes(app)

	// Start server
	err = app.Listen(":" + strconv.Itoa(AppConfig.AppPort))
	if err != nil {
		log.Fatalf("❌ Failed to start server %v", err)
	}
	fmt.Printf("🚀 Server is running on port %d\n", AppConfig.AppPort)
	return app
}
