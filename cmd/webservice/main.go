package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/PaulShpilsher/instalike/database"
	"github.com/PaulShpilsher/instalike/pkg/config"
	"github.com/PaulShpilsher/instalike/pkg/jwt"
	users "github.com/PaulShpilsher/instalike/pkg/users"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config := config.NewConfig()

	jwt.Initialize(&config.Jwt)

	db, _ := database.NewDbConnection(&config.Database)

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
