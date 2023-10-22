package users

import (
	"github.com/PaulShpilsher/instalike/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, serverConfig *config.ServerConfig, authMiddleware fiber.Handler, userService UserService) {

	usersRouter := router.Group("/users")

	// user sign up
	// POST: /users/register
	usersRouter.Post("/register", MakeUserRegisterHandler(userService))

	// user login
	// POST: /users/login
	usersRouter.Post("/login", MakeUserLoginHandler(serverConfig, userService))

	// get logged in user
	// GET: /users/me
	usersRouter.Get("/me", authMiddleware, MakeGetCurrentUserHandler(userService))
}
