package config

import "os"

type DatabaseConfig struct {
	ConnectionString string `json:"connection_string" split_words:"true"`
}

func databaseConfigFromEnvironment() *DatabaseConfig {
	var dbConfig DatabaseConfig
	dbConfig.ConnectionString = os.Getenv("DATABASE_CONNECTION_STRING")
	return &dbConfig
}

func NewDatabaseConfig() *DatabaseConfig {
	return databaseConfigFromEnvironment()
}
