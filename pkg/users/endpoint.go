package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type user struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func RegisterRoutes(router fiber.Router, s Service) {

	usersRouter := router.Group("/users")

	// POST: /users/signup
	usersRouter.Post(
		"/signup",

		// handler
		func(c *fiber.Ctx) error {
			dto := user{}
			if err := c.BodyParser(&dto); err != nil {
				return &fiber.Error{
					Code:    fiber.ErrBadRequest.Code,
					Message: err.Error(),
				}
			}
			log.Info(fmt.Sprintf("User signup.  Email: %s, Password %s", dto.Email, dto.Password))
			return c.JSON(dto)
		})

	usersRouter.Post("/login", func(c *fiber.Ctx) error {
		return c.SendString("login")
	})

}
