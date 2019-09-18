package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver for database/sql
)

// NewDatabase open connection to the database
func NewDatabase() *sqlx.DB {
	var (
		host     = getEnv("DB_HOST", "localhost")
		port     = getEnv("DB_PORT", "5432")
		user     = getEnv("DB_USER", "postgres")
		password = getEnv("DB_PASSWORD", "")
		dbname   = getEnv("DB_NAME", "todo-list")
		sslmode  = getEnv("DB_SSL_MODE", "disable") // verify-full
	)

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
