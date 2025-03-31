package loaders

import (
	"database/sql"
	"log"
	"server/config"
	"server/db/migrations"
)

func ConnectToDb() (*sql.DB, error) {
	AppConfig, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("❌ Failed to load config for app, error : %v", err)
	}

	dbAdminUser := migrations.AdminInformation{
		Host:     AppConfig.DbHost,
		Port:     AppConfig.DbPort,
		User:     AppConfig.DbAdminUser,
		Password: AppConfig.DbAdminPassword,
		Database: AppConfig.DbAdminDatabase,
	}

	connection, err := migrations.EnsureEnvironment(dbAdminUser, AppConfig.AppDbName,
		AppConfig.AppDbUser, AppConfig.AppDbPassword, migrations.Tables)

	if err != nil {
		log.Fatalf("❌ failed to ensure environment: %v", err)
	}

	log.Println("✅ Environment ensured successfully")

	return connection, nil

}
