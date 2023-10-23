package posts

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, authMiddleware fiber.Handler, s PostsService) {

	postsRouter := router.Group("/posts")
	postsRouter.Use(authMiddleware)

	postsRouter.Post("/", MakeCreatePostHandler(s))

	postsRouter.Get("/", MakeGetPostsHandler(s))

	postsRouter.Get("/:postId", MakeGetPostHandler(s))

	postsRouter.Put("/:postId", MakeUpdatePostHandler(s))

	postsRouter.Delete("/:postId", MakeDeletePostHandler(s))

	postsRouter.Post("/:postId/attachment", MakeUploadMediaFileToPostHandler(s))

	postsRouter.Post("/:postId/like", MakeLikePostHandler(s))

	postsRouter.Post("/:postId/comments", MakeCreatePostCommentHandler(s))

	postsRouter.Get("/:postId/comments", MakeGetPostCommentsHandler(s))

}
