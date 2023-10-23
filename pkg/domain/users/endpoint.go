package users

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/PaulShpilsher/instalike/pkg/config"
	"github.com/PaulShpilsher/instalike/pkg/middleware"
	"github.com/PaulShpilsher/instalike/pkg/utils"
)

func MakeUserRegisterHandler(s UsersService) fiber.Handler {
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
			if errors.Is(err, utils.ErrAlreadyExists) {
				return c.SendStatus(fiber.StatusConflict)
			} else {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		log.Debugf("user %s registered. id: %d", payload.Email, userId)
		return c.SendStatus(fiber.StatusCreated)
	}
}

func MakeUserLoginHandler(config *config.ServerConfig, s UsersService) fiber.Handler {
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

		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    token,
			Path:     "/",
			MaxAge:   int((time.Duration(config.TokenExpirationMinutes) * time.Minute).Seconds()),
			Secure:   false,
			HTTPOnly: true,
			Domain:   config.Domain,
		})

		return c.JSON(&loginOutput{
			Token: token,
		})
	}
}

func MakeUserMeHandler(s UsersService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		userId := middleware.GetAuthenicatedUserId(c)

		user, err := s.GetUserById(userId)
		if err != nil {
			log.Error(err)
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.JSON(&userOutput{
			UserId:  user.Id,
			Email:   user.Email,
			Created: user.Created,
			Updated: user.Updated,
		})
	}
}
