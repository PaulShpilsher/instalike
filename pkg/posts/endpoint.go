package posts

import (
	"errors"
	"strconv"
	"time"

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

		post, err := s.CreatePost(middleware.GetAuthenicatedUserId(c), payload.Contents)
		if err != nil {
			log.Error(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		log.Debugf("post created:  %v", post)
		return c.SendStatus(fiber.StatusCreated)
	}
}

///
/// get posts
///

type postOutput struct {
	Id        int       `json:"id"`
	Created   time.Time `json:"created"`
	IsUpdated bool      `json:"isUpdated"`

	UserId    int    `json:"userId"`
	Contents  string `json:"contents"`
	LikeCount int    `json:"likeCount"`
}

func makePostOutput(post Post) postOutput {
	return postOutput{
		Id:        post.Id,
		Created:   post.Created,
		IsUpdated: post.Created != post.Updated,
		Contents:  post.Contents,
		LikeCount: post.LikeCount,
	}
}

// MakeGetPostsHandler - get posts handler factory
func MakeGetPostsHandler(s PostsService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		posts, err := s.GetPosts()
		if err != nil {
			log.Error(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(utils.Map(posts, makePostOutput))
	}
}

// MakeGetPostByIdHandler - get post by id handler factory
func MakeGetPostByIdHandler(s PostsService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		postId, err := getPostId(c)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		post, err := s.GetPostById(postId)
		if err != nil {
			if errors.Is(err, utils.ErrNotFound) {
				return c.SendStatus(fiber.StatusNotFound)
			}

			log.Error(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(makePostOutput(post))
	}
}

// MakeDeletePostByIdHandler - get post by id handler factory
func MakeDeletePostByIdHandler(s PostsService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		postId, err := getPostId(c)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		err = s.DeletePostById(middleware.GetAuthenicatedUserId(c), postId)
		if err != nil {
			if errors.Is(err, utils.ErrNotFound) {
				return c.SendStatus(fiber.StatusNotFound)
			}

			if errors.Is(err, utils.ErrUnauthorized) {
				return c.SendStatus(fiber.StatusUnauthorized)
			}

			log.Error(err)
			return c.SendStatus(fiber.StatusInternalServerError)
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
