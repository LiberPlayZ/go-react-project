package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() (*Config, error) {
	var AppConfig *Config

	required := []string{
		"DB_HOST", "DB_PORT", "DB_ADMIN_USER", "DB_ADMIN_PASSWORD", "DB_ADMIN_DATABASE",
		"APP_DB_NAME", "APP_DB_USER", "APP_DB_PASSWORD", "APP_PORT",
	}

	err := godotenv.Load()
	log.Print(err)
	if err != nil {
		log.Fatalf("❌ Failed to load env file . ")
	}

	for _, key := range required {
		if os.Getenv(key) == "" {
			log.Fatalf("❌ ENV variable '%s' is required but not set.", key)
		}
	}

	port, err := GetIntEnv(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("❌ ENV DB_PORT must be a number: %v", err)

	}

	app_port, err := GetIntEnv(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatalf("❌ ENV DB_PORT must be a number: %v", err)

	}

	hashCost, err := GetIntEnv(os.Getenv("HASHING_COST"))
	if err != nil {
		log.Fatalf("❌ ENV DB_PORT must be a number: %v", err)

	}

	AppConfig = &Config{
		DbHost:          os.Getenv("DB_HOST"),
		DbPort:          port,
		DbAdminUser:     os.Getenv("DB_ADMIN_USER"),
		DbAdminPassword: os.Getenv("DB_ADMIN_PASSWORD"),
		DbAdminDatabase: os.Getenv("DB_ADMIN_DATABASE"),
		AppDbName:       os.Getenv("APP_DB_NAME"),
		AppDbUser:       os.Getenv("APP_DB_USER"),
		AppDbPassword:   os.Getenv("APP_DB_PASSWORD"),
		AppPort:         app_port,
		HashingCost:     hashCost,
	}

	log.Println("config load successfully")

	return AppConfig, nil

}
