package posts

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, authMiddleware fiber.Handler, s PostsService) {

	posts := router.Group("/posts")

	posts.Use(authMiddleware)

	// create new post
	// POST: /posts
	posts.Post("/register", notImplemented)

}

func notImplemented(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}
