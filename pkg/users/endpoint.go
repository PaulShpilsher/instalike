package users

import (
	"encoding/json"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/PaulShpilsher/instalike/pkg/middlewares"
	"github.com/PaulShpilsher/instalike/pkg/utils"
	"github.com/PaulShpilsher/instalike/pkg/utils/token"
)

type Service interface {
	Register(email string, password string) (userId int, err error)
	Login(email string, password string) (userId int, err error)
	GetUserById(id int) (user User, err error)
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
			return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorOutput(err.Error()))
		}

		if errors := utils.ValidateStruct(payload); errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.NewValidationErrorOutput(errors))
		}

		userId, err := s.Register(payload.Email, payload.Password)
		if err != nil {
			log.Error(err)
			// TODO: refactor to define specific business logic error for this.  For now use pgx error information
			if strings.Contains(err.Error(), "duplicate key value violates unique") {
				return c.Status(fiber.StatusConflict).JSON(utils.NewErrorOutput("User with that email already exists"))
			} else {
				return c.Status(fiber.StatusInternalServerError).JSON(utils.NewErrorOutput("Server error"))
			}
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
			return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorOutput(err.Error()))
		}

		if errors := utils.ValidateStruct(payload); errors != nil {
			err, _ := json.Marshal(errors)
			log.Error(string(err))
			return c.SendStatus(fiber.StatusUnauthorized) // do not give to the client hints what is wrong with login data
		}

		userId, err := s.Login(payload.Email, payload.Password)
		if err != nil {
			log.Error(err)
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Greate JWT
		token, err := token.CreateJwtToken(userId)
		if err != nil {
			log.Error(err)
			return c.Status(fiber.StatusInternalServerError).JSON(utils.NewErrorOutput(err.Error()))
		}

		log.Debugf("user %s logged in. id: %d", payload.Email, userId)
		return c.JSON(&loginOutput{
			UserId: userId,
			Token:  token,
		})
	}

	// me - get logged in user information
	// IMPORTANT: this is a test function, remove it before production
	me := func(c *fiber.Ctx) error {

		userId := middlewares.GetAuthenicatedUserId(c)
		user, err := s.GetUserById(userId)
		if err != nil {
			log.Error(err)
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.JSON(fiber.Map{
			"email":   user.Email,
			"created": user.CreatedAt,
			"updated": user.UpdatedAt,
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
	usersRouter.Get("/me", middlewares.AuthenticateUser, me)
}
