package config

import (
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Schema   string
	SSLMode  string
}

func GetDBConfig() *DBConfig {
	config := DBConfig{}
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "production" {
		config.Host = os.Getenv("DB_HOST")
		config.Port = os.Getenv("DB_PORT")
		config.User = os.Getenv("DB_USER")
		config.Password = os.Getenv("DB_PASSWORD")
		config.Schema = os.Getenv("DB_NAME")
		config.SSLMode = "enable"
	} else {
		config.Host = "localhost"
		config.Port = "5432"
		config.User = "postgres"
		config.Password = "postgres"
		config.Schema = "postgres"
		config.SSLMode = "disable"
	}

	return &config
}

type ServerConfig struct {
	Address string
	Port    string
}

func GetServerConfig() *ServerConfig {
	return &ServerConfig{
		Address: "127.0.0.1",
		Port:    "8010",
	}
}
