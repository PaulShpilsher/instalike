package posts

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/PaulShpilsher/instalike/pkg/middleware"
	"github.com/PaulShpilsher/instalike/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

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
func MakeCreatePostHandler(s PostsService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var payload createPostInput
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorOutput(err.Error()))
		}

		if errors := utils.ValidateStruct(payload); errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.NewValidationErrorOutput(errors))
		}

		postId, err := s.CreatePost(middleware.GetAuthenicatedUserId(c), payload.Body)
		if err != nil {
			log.Error(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(fiber.StatusCreated).JSON(createPostOutput{
			Id: postId,
		})
	}
}

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
func MakeGetPostsHandler(s PostsService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		posts, err := s.GetPosts()
		if err != nil {
			log.Error(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(utils.Map(posts, makeGetPostOutput))
	}
}

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
func MakeGetPostHandler(s PostsService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		postId, err := getPostId(c)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		post, err := s.GetPost(postId)
		if err != nil {
			if errors.Is(err, utils.ErrNotFound) {
				return c.SendStatus(fiber.StatusNotFound)
			}

			log.Error(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(makeGetPostOutput(post))
	}
}

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
func MakeDeletePostHandler(s PostsService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		postId, err := getPostId(c)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		err = s.DeletePost(middleware.GetAuthenicatedUserId(c), postId)
		if err != nil {
			if errors.Is(err, utils.ErrNotFound) {
				return c.SendStatus(fiber.StatusNotFound)
			}

			if errors.Is(err, utils.ErrForbidden) {
				return c.SendStatus(fiber.StatusForbidden)
			}

			log.Error(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

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
func MakeUpdatePostHandler(s PostsService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		postId, err := getPostId(c)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		var payload updatePostInput
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorOutput(err.Error()))
		}

		if errors := utils.ValidateStruct(payload); errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorOutput(err.Error()))
		}

		err = s.UpdatePost(middleware.GetAuthenicatedUserId(c), postId, payload.Body)
		if err != nil {
			if errors.Is(err, utils.ErrNotFound) {
				return c.SendStatus(fiber.StatusNotFound)
			}

			if errors.Is(err, utils.ErrForbidden) {
				return c.SendStatus(fiber.StatusForbidden)
			}

			log.Error(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

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
func MakeUploadMediaFileToPostHandler(s PostsService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		postId, err := getPostId(c)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		// TODO: put it in config
		const MaxUploadSize = 1024 * 1024 * 10 // 10Mb
		if file.Size > MaxUploadSize {
			return c.SendStatus(fiber.StatusRequestEntityTooLarge)
		} else if file.Size == 0 {
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		}

		contentType := file.Header["Content-Type"][0]

		if matched, err := regexp.MatchString("^(:?image|video)/\\w{2,}$", contentType); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		} else if !matched {
			return c.Status(fiber.StatusBadRequest).SendString("Only images and videos are allowed")
		}

		// get buffer from file
		buffer, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		defer buffer.Close()

		userId := middleware.GetAuthenicatedUserId(c)
		err = s.AttachFileToPost(userId, postId, contentType, int(file.Size), buffer)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

func getPostId(c *fiber.Ctx) (int, error) {
	param := c.Params("postId")
	value, err := strconv.Atoi(param)
	if err != nil {
		log.Errorf("bad post id param: %s. err: %v", param, err)
		return 0, err
	}
	return value, nil
}
