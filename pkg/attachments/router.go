package attachments

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, authMiddleware fiber.Handler, attachmentsService AttachmentsService) {

	files := router.Group("/attachments")
	// files.Use(authMiddleware)

	// download attachments
	files.Get("/:attachmentId", MakeCreatePostHandler(attachmentsService))
}
