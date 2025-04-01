package loaders

import (
	"database/sql"
	"log"
	"server/config"
	"server/db/migrations"
)

func ConnectToDb(AppConfig *config.Config) (*sql.DB, error) {

	dbAdminUser := migrations.AdminInformation{
		Host:     AppConfig.DbHost,
		Port:     AppConfig.DbPort,
		User:     AppConfig.DbAdminUser,
		Password: AppConfig.DbAdminPassword,
		Database: AppConfig.DbAdminDatabase,
	}

	log.Println("Starting connecting to db")

	connection, err := migrations.EnsureEnvironment(dbAdminUser, AppConfig.AppDbName,
		AppConfig.AppDbUser, AppConfig.AppDbPassword, migrations.Tables)

	if err != nil {
		log.Fatalf("❌ failed to ensure environment: %v", err)
	}

	log.Println("✅ Environment ensured successfully")
	log.Println("✅ Db connected successfully")

	return connection, nil

}
