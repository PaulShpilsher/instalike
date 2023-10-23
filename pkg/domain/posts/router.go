package posts

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, authMiddleware fiber.Handler, s PostsService) {

	postsRouter := router.Group("/posts")
	postsRouter.Use(authMiddleware)

	// CreatePost godoc
	// @Summary Create post
	// @Description creates a new post
	// @Tags posts
	// @Security Bearer
	// @param Authorization header string true "Authorization"
	// @Accept json
	// @Produce json
	// @Param data body createPostInput true "The input post struct"
	// @Success 200 {object} createPostOutput
	// @Failure 400
	// @Failure 401
	// @Router /api/posts [post]
	postsRouter.Post("/", MakeCreatePostHandler(s))

	// GetPosts godoc
	// @Summary Gets all posts
	// @Description gets all posts
	// @Tags posts
	// @Security Bearer
	// @param Authorization header string true "Authorization"
	// @Produce json
	// @Success 200 {array} getPostOutput
	// @Failure 400
	// @Failure 401
	// @Router /api/posts [get]
	postsRouter.Get("/", MakeGetPostsHandler(s))

	// GetPost godoc
	// @Summary Gets a post by post id
	// @Description gets a post by post id
	// @Tags posts
	// @Security Bearer
	// @param Authorization header string true "Authorization"
	// @Produce json
	// @Param postId path int true "Post ID"
	// @Success 200 {object} getPostOutput
	// @Failure 400
	// @Failure 401
	// @Failure 404
	// @Router /api/posts/{postId} [get]
	postsRouter.Get("/:postId", MakeGetPostHandler(s))

	// UpdatePost godoc
	// @Summary Updates post
	// @Description updates post. only author of the post is allowed to update the post
	// @Tags posts
	// @Security Bearer
	// @param Authorization header string true "Authorization"
	// @Accept json
	// @Param postId path int true "Post ID"
	// @Param data body updatePostInput true "The update post struct"
	// @Success 204
	// @Failure 400
	// @Failure 401
	// @Failure 403
	// @Failure 404
	// @Router /api/posts/{postId} [put]
	postsRouter.Put("/:postId", MakeUpdatePostHandler(s))

	// DeletePost godoc
	// @Summary Deletes post
	// @Description deletes post by post id. only author of the post is allowed to delete the post
	// @Tags posts
	// @Security Bearer
	// @param Authorization header string true "Authorization"
	// @Param postId path int true "Post ID"
	// @Success 204
	// @Failure 400
	// @Failure 401
	// @Failure 403
	// @Failure 404
	// @Router /api/posts/{postId} [delete]
	postsRouter.Delete("/:postId", MakeDeletePostHandler(s))

	// UploadPostAttachment godoc
	// @Summary Attaches multimedia file to post
	// @Description attaches multimedia file to post. only author of the post is allowed to add files to the post. only images and videos are allowed to be uploaded
	// @Tags posts
	// @Security Bearer
	// @param Authorization header string true "Authorization"
	// @Accept mpfd
	// @Param postId path int true "Post ID"
	// @Param file formData file true "The file upload form"
	// @Success 204
	// @Failure 400
	// @Failure 401
	// @Failure 403
	// @Failure 404
	// @Failure 413
	// @Failure 422
	// @Router /api/posts/{postId}/attachment [post]
	postsRouter.Post("/:postId/attachment", MakeUploadMediaFileToPostHandler(s))

}
