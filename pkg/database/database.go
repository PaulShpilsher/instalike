package database

import (
	"log"
	"time"

	"github.com/PaulShpilsher/instalike/pkg/config"
	_ "github.com/jackc/pgx/v5/stdlib" // Standard library bindings for pgx
	"github.com/jmoiron/sqlx"
)

func NewDbConnection(config *config.DatabaseConfig) *sqlx.DB {

	// Define database connection for PostgreSQL.
	// TODO: add retry with timeout logic
	db, err := sqlx.Connect("pgx", config.Url)
	if err != nil {
		log.Panicf("not connected to database, err: %v", err)
	}

	// Set database connection settings.
	db.SetMaxOpenConns(config.MaxOpenConnections)                      // the default is 0 (unlimited)
	db.SetMaxIdleConns(config.MaxIdleConnections)                      // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(time.Duration(config.MaxLifetimeConnctions)) // 0, connections are reused forever

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close()
		log.Panicf("failed to ping to database, %v", err)
	}

	return db
}
