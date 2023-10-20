package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/PaulShpilsher/instalike/database/postgres"
	users "github.com/PaulShpilsher/instalike/pkg/users"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	loadConfig()

	db, _ := postgres.PostgreSQLConnection()

	// Create API server
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	apiRoute := app.Group("/api")

	// /api/users

	// Service layer

	repository := users.NewRepository(db)

	service := users.NewService(repository)

	// Endpoint layer
	users.RegisterRoutes(apiRoute, service)

	// Starting server with graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := app.Listen(os.Getenv("SERVER_URL")); err != nil {
			log.Fatalf("listen returned err: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("got interruption signal")
	if err := app.ShutdownWithContext(context.TODO()); err != nil {
		log.Printf("shutdown returned an err: %v\n", err)
	}

	db.Close()

	log.Println("done")

}

func loadConfig() {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	envFile := filepath.Join(filepath.Dir(ex), ".env")

	if err := godotenv.Load(envFile); err != nil {
		panic("Error loading .env file")
	}
}
