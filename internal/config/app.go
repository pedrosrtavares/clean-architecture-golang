package config

import "os"

type AppConfig struct {
	Environment string `json:"environment" split_words:"true"`
	Port        string `json:"port" split_words:"true"`
	DBConfig    *DatabaseConfig
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Environment: os.Getenv("ENVIRONMENT"),
		Port:        os.Getenv("PORT"),
		DBConfig:    NewDatabaseConfig(),
	}
}
