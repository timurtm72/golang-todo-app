package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	} else {
		fmt.Printf("Database: %s connection established\n", "postgres")
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	} else {
		fmt.Printf("Database: %s connection pinged\n", "postgres")
	}
	return db, nil
}
