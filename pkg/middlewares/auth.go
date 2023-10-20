package middlewares

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/PaulShpilsher/instalike/pkg/utils/token"
)

func AuthenticateUser(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims, err := token.ValidateJwtToken(tokenString)
	if err != nil {
		log.Error(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	jti := claims["jti"].(string)
	userId, err := strconv.Atoi(jti)
	if err != nil {
		log.Error(err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	c.Locals("userId", userId)
	c.Next()
	return nil
}
