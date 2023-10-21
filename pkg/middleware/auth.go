package middleware

import (
	"strconv"
	"strings"

	"github.com/PaulShpilsher/instalike/pkg/token"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func GetAuthMiddleware(s token.JwtService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		token := getToken(c)
		if token == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// validate authorization token
		data, err := s.ValidateToken(token)
		if err != nil {
			log.Error(err)
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		userId, err := strconv.Atoi(data)
		if err != nil {
			log.Error(err)
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Locals("userId", userId)
		c.Next()
		return nil
	}
}

func getToken(c *fiber.Ctx) string {
	// get authorization token
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		return strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		return c.Cookies("token")
	} else {
		return authorization
	}
}

func GetAuthenicatedUserId(c *fiber.Ctx) int {
	userId := c.Locals("userId").(int)
	return userId
}
