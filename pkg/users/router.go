package users

import (
	"github.com/gofiber/fiber/v2"

	"github.com/PaulShpilsher/instalike/pkg/middlewares"
)

func RegisterRoutes(router fiber.Router, s Service) {

	usersRouter := router.Group("/users")

	// user sign up
	// POST: /users/register
	usersRouter.Post("/register", MakeUserRegisterHandler(s))

	// user login
	// POST: /users/login
	usersRouter.Post("/login", MakeUserLoginHandler(s))

	// get logged in user
	// GET: /users/me
	usersRouter.Get("/me", middlewares.AuthenticateUser, MakeGetLoggedInUserHandler(s))
}
