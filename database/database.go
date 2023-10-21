package database

import (
	"fmt"
	"time"

	"github.com/PaulShpilsher/instalike/pkg/config"
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/jackc/pgx/v5/stdlib" // Standard library bindings for pgx
	"github.com/jmoiron/sqlx"
)

func NewDbConnection(config *config.DatabaseConfig) (*sqlx.DB, error) {

	// Define database connection for PostgreSQL.
	db, err := sqlx.Connect("pgx", config.Url)
	if err != nil {
		log.Fatalf("error, not connected to database, %v", err)
	}

	// Set database connection settings.
	db.SetMaxOpenConns(config.MaxOpenConnections)                      // the default is 0 (unlimited)
	db.SetMaxIdleConns(config.MaxIdleConnections)                      // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(time.Duration(config.MaxLifetimeConnctions)) // 0, connections are reused forever

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}
