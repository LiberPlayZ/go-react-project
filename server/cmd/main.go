package main

import (
	"log"
	"server/config"
	"server/db/loaders"
	"server/internal/api/server"
)

func main() {
	// load config
	AppConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load config for app, error : %v", err)
	}

	db, err := loaders.ConnectToDb(AppConfig)
	if err != nil {
		log.Fatalf("❌ Could not connect to database: %v", err)
	}

	defer db.Close()

	server.NewServer(AppConfig.AppPort)

}
