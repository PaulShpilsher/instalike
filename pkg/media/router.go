package media

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(mediaRouter fiber.Router, authMiddleware fiber.Handler, mediaService MediaService) {

	mediaRouter.Use(authMiddleware)

	// download attachments
	mediaRouter.Get("/attachments/:attachmentId", MakeDownlodPostAttachmentHandler(mediaService))
}
