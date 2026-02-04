package config

import (
	"fmt"
	"os"
)

// Config holds all configuration for the application.
type Config struct {
	Port string
	DB   DatabaseConfig
}

// DatabaseConfig holds all database-related configuration.
type DatabaseConfig struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
}

// DSN returns the Data Source Name for connecting to PostgreSQL.
func (db DatabaseConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		db.User, db.Password, db.Host, db.Port, db.Name)
}

// Load initializes the configuration from environment variables.
func Load() *Config {
	return &Config{
		Port: getEnv("PORT", "8080"),
		DB: DatabaseConfig{
			User:     getEnv("DB_USER", "gopher"),
			Password: getEnv("DB_PASSWORD", "password"),
			Name:     getEnv("DB_NAME", "gophership"),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
		},
	}
}

// getEnv retrieves a value from the environment or returns the fallback.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
