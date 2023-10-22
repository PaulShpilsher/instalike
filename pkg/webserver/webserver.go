package webserver

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/PaulShpilsher/instalike/pkg/attachments"
	"github.com/PaulShpilsher/instalike/pkg/config"
	"github.com/PaulShpilsher/instalike/pkg/database"
	"github.com/PaulShpilsher/instalike/pkg/middleware"
	"github.com/PaulShpilsher/instalike/pkg/posts"
	"github.com/PaulShpilsher/instalike/pkg/token"
	"github.com/PaulShpilsher/instalike/pkg/users"
	"github.com/jmoiron/sqlx"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type WebServer struct {
	serverAddress string
	app           *fiber.App
	db            *sqlx.DB
}

func NewWebServer(config *config.Config) WebServer {

	db := database.NewDbConnection(&config.Database)

	jwtService := token.NewJwtService(&config.Server)
	authMiddleware := middleware.GetAuthMiddleware(jwtService)

	app := fiber.New()
	app.Use(cors.New())

	api := fiber.New()

	app.Static(
		"/static",  // mount address
		"./public", // path to the file folder
	)

	app.Mount("/api", api)

	// /api/users
	{
		usersRepository := users.NewRepository(db)
		usersService := users.NewService(usersRepository, jwtService)
		users.RegisterRoutes(api, &config.Server, authMiddleware, usersService)
	}

	// /api/posts
	{
		postsRepository := posts.NewPostsRepository(db)
		attachmentRepository := posts.NewAttachmentRepository(db)
		postsService := posts.NewPostsService(postsRepository, attachmentRepository)
		posts.RegisterRoutes(api, authMiddleware, postsService)
	}

	filer := fiber.New()
	app.Mount("/filer", filer)

	// /filer/attachments
	{
		attachmentsRepository := attachments.NewAttachmentRepository(db)
		attachmentsService := attachments.NewAttachmentService(attachmentsRepository)
		attachments.RegisterRoutes(filer, authMiddleware, attachmentsService)
	}

	return WebServer{
		app:           app,
		db:            db,
		serverAddress: fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
	}
}

// Start function for starting server with a graceful shutdown.
func (s WebServer) Start() {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
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
	if err := s.app.Listen(s.serverAddress); err != nil {
		log.Fatalf("server failed to start. err : %v", err)
	}

	<-idleConnsClosed
}
