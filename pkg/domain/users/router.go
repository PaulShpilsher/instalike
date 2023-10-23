package users

import (
	"github.com/PaulShpilsher/instalike/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(apiRouter fiber.Router, serverConfig *config.ServerConfig, authMiddleware fiber.Handler, userService UsersService) {

	usersRouter := apiRouter.Group("/users")

	usersRouter.Post("/register", MakeUserRegisterHandler(userService))

	usersRouter.Post("/login", MakeUserLoginHandler(serverConfig, userService))

	usersRouter.Get("/me", authMiddleware, MakeUserMeHandler(userService)) // NOTE: In future move this to users domain
}
