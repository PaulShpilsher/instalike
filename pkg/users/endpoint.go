package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Service interface {
	Signup(email string, password string) (userId int, token string, err error)
	Login(email string, password string) (userId int, token string, err error)
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

	// POST: /users/signup
	usersRouter.Post(
		"/signup",

		// handler
		func(c *fiber.Ctx) error {
			user := userRequest{}
			if err := c.BodyParser(&user); err != nil {
				return &fiber.Error{
					Code:    fiber.ErrBadRequest.Code,
					Message: err.Error(),
				}
			}

			userId, token, err := s.Signup(user.Email, user.Password)
			if err != nil {
				log.Error(err.Error())
				return err
			}

			log.Info(fmt.Sprintf("user signed up. id(%d), token(%s)", userId, token))
			return c.JSON(&userResponse{
				UserId: userId,
				Token:  token,
			})
		})

	usersRouter.Post("/login", func(c *fiber.Ctx) error {
		return c.SendString("login")
	})

}
