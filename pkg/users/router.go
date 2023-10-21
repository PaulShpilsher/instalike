package users

import (
	"github.com/PaulShpilsher/instalike/pkg/token"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, authMiddleware fiber.Handler, s UserService, j token.JwtService) {

	usersRouter := router.Group("/users")

	// user sign up
	// POST: /users/register
	usersRouter.Post("/register", MakeUserRegisterHandler(s))

	// user login
	// POST: /users/login
	usersRouter.Post("/login", MakeUserLoginHandler(s, j))

	// get logged in user
	// GET: /users/current
	usersRouter.Get("/current", authMiddleware, MakeGetCurrentUserHandler(s))
}
