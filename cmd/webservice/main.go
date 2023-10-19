package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	users "instalike/pkg/users"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUri := os.Getenv("DB_URI")
	fmt.Printf("DB_URI=%s", dbUri)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	apiRoute := app.Group("/api")

	service := users.Service{}

	users.RegisterRoutes(apiRoute, service)

	// Starting server with graceful shutdown

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := app.Listen(os.Getenv("HOST")); err != nil {
			log.Fatalf("listen returned err: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("got interruption signal")
	if err := app.ShutdownWithContext(context.TODO()); err != nil {
		log.Printf("shutdown returned an err: %v\n", err)
	}

	log.Println("done")

}
