package users

import (
	"github.com/PaulShpilsher/instalike/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(apiRouter fiber.Router, serverConfig *config.ServerConfig, authMiddleware fiber.Handler, userService UsersService) {

	usersRouter := apiRouter.Group("/users")

	// Register doc
	// @Summary User register
	// @Description registers user.
	// @Tags auth
	// @Accept json
	// @Produce json
	// @Param data body registerInput true "The input register struct"
	// @Success 201
	// @Failure 400 {object} utils.ErrorOutput
	// @Failure 409
	// @Router /api/users/register [post]
	usersRouter.Post("/register", MakeUserRegisterHandler(userService))

	// Login godoc
	// @Summary User login
	// @Description performs user login, returns jwt token and sets http-only cookie.
	// @Tags auth
	// @Accept json
	// @Produce json
	// @Param data body loginInput true "The input login struct"
	// @Success 200 {object} loginOutput
	// @Failure 400 {object} utils.ErrorOutput
	// @Failure 401
	// @Router /api/users/login [post]
	usersRouter.Post("/login", MakeUserLoginHandler(serverConfig, userService))

	// Me godoc
	// @Summary Current user information
	// @Description gets currenly logger in user information
	// @Tags auth
	// @Security Bearer
	// @param Authorization header string true "Authorization"
	// @Produce json
	// @Success 200 {object} userOutput
	// @Failure 401
	// @Failure 404
	// @Router /api/users/me [get]
	usersRouter.Get("/me", authMiddleware, MakeUserMeHandler(userService)) // NOTE: In future move this to users domain
}
