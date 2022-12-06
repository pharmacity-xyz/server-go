package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "password",
		Database: "pharmacitydb",
		SSLMode:  "disable",
	}
}

func Open() (*sql.DB, error) {
	mode := os.Getenv("MODE")
	var dbString string

	if mode == "PRODUCTION" {
		dbString = os.Getenv("DB_STRING_PRODUCTION")
	} else {
		dbString = os.Getenv("DB_STRING_LOCAL")
	}

	db, err := sql.Open("pgx", dbString)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	return db, nil
}
