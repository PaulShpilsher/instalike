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

///
/// Create post
///

// MakeCreatePostHandler - create post handler factory
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

///
/// get posts
///

// MakeGetPostsHandler - get posts handler factory
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

///
/// get post
///

// MakeGetPostHandler - get post by id handler factory
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

///
/// delete post
///

// MakeDeletePostHandler - delete post by id handler factory
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

///
/// update post
///

// MakeUpdatePostHandler - update post by id handler factory
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

///
/// Upload files
///

const MaxUploadSize = 1024 * 1024 // 1Mb

// MakeUploadMediaFileToPostHandler - upload multimedia files
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
