package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/PaulShpilsher/instalike/pkg/middlewares"
	"github.com/PaulShpilsher/instalike/pkg/utils"
	"github.com/PaulShpilsher/instalike/pkg/utils/token"
)

type Service interface {
	Register(email string, password string) (userId int, err error)
	Login(email string, password string) (userId int, err error)
}

type registerInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=5"`
}

type registerOutput struct {
	UserId int `json:"userId"`
}

type loginInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type loginOutput struct {
	UserId int    `json:"userId"`
	Token  string `json:"token"`
}

func RegisterRoutes(router fiber.Router, s Service) {

	// API handlers

	// register - registers new user
	register := func(c *fiber.Ctx) error {

		var payload registerInput
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}

		if errors := utils.ValidateStruct(payload); errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
		}

		userId, err := s.Register(payload.Email, payload.Password)
		if err != nil {
			log.Error(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}

		log.Debugf("user %s registered. id: %d", payload.Email, userId)
		return c.Status(fiber.StatusCreated).JSON(registerOutput{
			UserId: userId,
		})
	}

	// login - user login
	login := func(c *fiber.Ctx) error {
		var payload loginInput
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}

		userId, err := s.Login(payload.Email, payload.Password)
		if err != nil {
			log.Error(err)
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Greate JWT
		token, err := token.CreateJwtToken(userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}

		log.Debugf("user %s logged in. id: %d", payload.Email, userId)
		return c.JSON(&loginOutput{
			UserId: userId,
			Token:  token,
		})
	}

	// regiser handlers

	usersRouter := router.Group("/users")

	// POST: /users/register
	usersRouter.Post("/register", register)

	// POST: /users/login
	usersRouter.Post("/login", login)

	// TEST
	// Gets logged in user information
	// GET: /users/me
	usersRouter.Get("/me", middlewares.AuthenticateUser, func(c *fiber.Ctx) error {
		userId := middlewares.GetAuthenicatedUserId(c)
		return c.SendString(fmt.Sprintf("userId: %d", userId))
	})
}
