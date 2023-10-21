package posts

import (
	"strconv"

	"github.com/PaulShpilsher/instalike/pkg/middleware"
	"github.com/PaulShpilsher/instalike/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

///
/// Create post
///

type createPostInput struct {
	Contents string `json:"contents" validate:"required"`
}

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

		userId := middleware.GetAuthenicatedUserId(c)

		post, err := s.CreatePost(userId, payload.Contents)
		if err != nil {
			log.Error(err)
			return c.Status(fiber.StatusInternalServerError).JSON(utils.NewErrorOutput("Server error"))
		}

		log.Debugf("post created:  %v", post)
		return c.SendStatus(fiber.StatusCreated)
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
