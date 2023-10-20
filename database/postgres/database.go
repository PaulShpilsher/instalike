package postgres

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2/log"
	_ "github.com/jackc/pgx/v5/stdlib" // Standard library bindings for pgx
	"github.com/jmoiron/sqlx"
)

func PostgreSQLConnection() (*sqlx.DB, error) {

	// Define database connection for PostgreSQL.
	url := os.Getenv("DB_URL")
	db, err := sqlx.Connect("pgx", url)
	if err != nil {
		log.Fatalf("error, not connected to database, %v", err)
	}

	// Define database connection settings.
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// Set database connection settings.
	db.SetMaxOpenConns(maxConn)                           // the default is 0 (unlimited)
	db.SetMaxIdleConns(maxIdleConn)                       // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn)) // 0, connections are reused forever

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}
