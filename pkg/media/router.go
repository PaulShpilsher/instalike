package media

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(mediaRouter fiber.Router, authMiddleware fiber.Handler, attachmentsService AttachmentsService) {

	mediaRouter.Use(authMiddleware)

	// download attachments
	mediaRouter.Get("/attachments/:attachmentId", MakeCreatePostHandler(attachmentsService))
}
