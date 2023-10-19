package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Service interface {
	Signup(email string, password string) (userId int, err error)
	Login(email string, password string) (userId int, err error)
}

type userRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type userResponse struct {
	UserId int    `json:"userId"`
	Token  string `json:"token"`
}

func RegisterRoutes(router fiber.Router, s Service) {

	usersRouter := router.Group("/users")

	// POST: /users/register
	usersRouter.Post("/register", func(c *fiber.Ctx) error {
		user, err := parseUserRequest(c)
		if err != nil {
			log.Error(err.Error())
			return err
		}

		userId, err := s.Login(user.Email, user.Password)
		if err != nil {
			log.Error(err.Error())
			return err
		}

		log.Info(fmt.Sprintf("user signed up. id(%d)", userId))
		return c.SendStatus(fiber.StatusCreated)
	})

	// POST: /users/login
	usersRouter.Post("/login", func(c *fiber.Ctx) error {
		user, err := parseUserRequest(c)
		if err != nil {
			return err
		}

		userId, err := s.Login(user.Email, user.Password)
		if err != nil {
			log.Error(err.Error())
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// TODO: Greate JWT
		token := "authToken"

		log.Info(fmt.Sprintf("user logged in. id(%d), token(%s)", userId, token))
		return c.JSON(&userResponse{
			UserId: userId,
			Token:  token,
		})
	})
}

func parseUserRequest(c *fiber.Ctx) (*userRequest, error) {
	user := userRequest{}
	if err := c.BodyParser(&user); err != nil {
		log.Error(err)
		return nil, &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}
	return &user, nil
}
