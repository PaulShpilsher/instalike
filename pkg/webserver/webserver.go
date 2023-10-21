package webserver

import (
	"log"
	"os"
	"os/signal"

	"github.com/PaulShpilsher/instalike/database"
	"github.com/PaulShpilsher/instalike/pkg/config"
	"github.com/PaulShpilsher/instalike/pkg/middleware"
	"github.com/PaulShpilsher/instalike/pkg/token"
	"github.com/PaulShpilsher/instalike/pkg/users"
	"github.com/jmoiron/sqlx"

	"github.com/gofiber/fiber/v2"
)

type WebServer struct {
	serverUrl string
	app       *fiber.App
	db        *sqlx.DB
}

func NewWebServer(config *config.Config) WebServer {

	db := database.NewDbConnection(&config.Database)

	jwtService := token.NewJwtService(&config.Jwt)

	app := fiber.New()
	authMiddleware := middleware.GetAuthMiddleware(jwtService)

	api := fiber.New()
	app.Mount("/api", api)

	// /api/users
	{
		repository := users.NewRepository(db)
		service := users.NewService(repository, jwtService)
		users.RegisterRoutes(api, authMiddleware, service)
	}

	return WebServer{
		app:       app,
		db:        db,
		serverUrl: config.Server.Url,
	}
}

// Start function for starting server with a graceful shutdown.
func (s WebServer) Start() {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := s.app.Shutdown(); err != nil {
			log.Printf("server is not shutting down. err: %v", err)
		} else {
			log.Println("server is shutdown")
		}

		if err := s.db.Close(); err != nil {
			log.Printf("database connection is not closing. err: %v", err)
		} else {
			log.Println("database connection is closed")
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := s.app.Listen(s.serverUrl); err != nil {
		log.Fatalf("server failed to start. err : %w", err)
	}

	<-idleConnsClosed
}
