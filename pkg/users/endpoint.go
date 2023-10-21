package users

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/PaulShpilsher/instalike/pkg/middleware"
	"github.com/PaulShpilsher/instalike/pkg/utils"
)

// Service layer interface
type Service interface {
	Register(email string, password string) (userId int, err error)
	Login(email string, password string) (userId int, token string, err error)
	GetUserById(id int) (user User, err error)
}

///
/// Sign up
///

type registerInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=5"`
}

// MakeUserRegisterHandler - register handler factory
func MakeUserRegisterHandler(s Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

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
			if strings.Contains(err.Error(), "user already exists") {
				return c.Status(fiber.StatusConflict).JSON(utils.NewErrorOutput("user already exists"))
			} else {
				return c.Status(fiber.StatusInternalServerError).JSON(utils.NewErrorOutput("Server error"))
			}
		}

		log.Debugf("user %s registered. id: %d", payload.Email, userId)
		return c.SendStatus(fiber.StatusCreated)
	}
}

///
/// Login
///

type loginInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type loginOutput struct {
	UserId int    `json:"userId"`
	Token  string `json:"token"`
}

// MakeUserLoginHandler - login handler factory
func MakeUserLoginHandler(s Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload loginInput
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(utils.NewErrorOutput(err.Error()))
		}

		if errors := utils.ValidateStruct(payload); errors != nil {
			err, _ := json.Marshal(errors)
			log.Error(string(err))
			return c.SendStatus(fiber.StatusUnauthorized) // do not give to the client hints what is wrong with login data
		}

		userId, token, err := s.Login(payload.Email, payload.Password)
		if err != nil {
			log.Error(err)
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		log.Debugf("user %s logged in. id: %d", payload.Email, userId)
		return c.JSON(&loginOutput{
			UserId: userId,
			Token:  token,
		})
	}
}

///
/// Current user information
///

type currentUserOutput struct {
	UserId  int       `json:"userId"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// MakeGetCurrentUserHandler - get logged in user information
func MakeGetCurrentUserHandler(s Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		userId := middleware.GetAuthenicatedUserId(c)

		user, err := s.GetUserById(userId)
		if err != nil {
			log.Error(err)
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.JSON(&currentUserOutput{
			UserId:  user.Id,
			Email:   user.Email,
			Created: user.Created,
			Updated: user.Updated,
		})
	}
}
