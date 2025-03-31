package loaders

import (
	"log"
	"server/config"
	migrations "server/db/migration"
)

func ConnectToDb() {
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

	defer connection.Close()
	log.Println("✅ Environment ensured successfully")

}
