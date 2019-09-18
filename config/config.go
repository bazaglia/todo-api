package config

import (
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Enabled  bool
	Database *sqlx.DB
	Port     string
}

func NewConfig(db *sqlx.DB) *Config {
	return &Config{
		Enabled:  true,
		Database: db,
		Port:     "8000",
	}
}
