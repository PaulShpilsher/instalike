package media

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(mediaRouter fiber.Router, authMiddleware fiber.Handler, mediaService MediaService) {

	mediaRouter.Get("/media/attachments/:attachmentId", authMiddleware, MakeDownlodPostAttachmentHandler(mediaService))

}
