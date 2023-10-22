package posts

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router, authMiddleware fiber.Handler, s PostsService) {

	posts := api.Group("/posts")
	posts.Use(authMiddleware)

	// Create a new post (POST /api/posts)
	posts.Post("/", MakeCreatePostHandler(s))

	// List all posts (GET /api/posts).
	posts.Get("/", MakeGetPostsHandler(s))

	// get a specific post by ID (GET /api/posts/{postID}).
	posts.Get("/:postId", MakeGetPostHandler(s))

	// Update an existing post (PUT /api/posts/{postID}).
	posts.Put("/:postId", MakeUpdatePostHandler(s))

	// Delete a post (DELETE /api/posts/{postID}).
	posts.Delete("/:postId", MakeDeletePostHandler(s))

	// Add multimedia file to post (POST /api/posts/{postID}/file)
	posts.Post("/:postId/attachment", MakeUploadMediaFileToPostHandler(s))
	// temp := api.Group("/multimedia")
	// temp.Post("/:postId/upload", MakeUploadMediaFileToPostHandler(s))

}

// func notImplemented(c *fiber.Ctx) error {
// 	return c.SendStatus(fiber.StatusNotImplemented)
// }
