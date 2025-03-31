package config

type Config struct {
	DbHost          string
	DbPort          int
	DbAdminUser     string
	DbAdminPassword string
	DbAdminDatabase string
	AppDbName       string
	AppDbUser       string
	AppDbPassword   string
}
