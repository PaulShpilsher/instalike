package media

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(mediaRouter fiber.Router, authMiddleware fiber.Handler, mediaService MediaService) {

	mediaRouter.Use(authMiddleware)

	// DownloadPostAMultimediAttachment godoc
	// @Summary Downloads multimedia attached to a post
	// @Description downloads multimedia attached to a post
	// @Tags media
	// @Security Bearer
	// @param Authorization header string true "Authorization"
	// @Produce */*
	// @Param attachmentId path int true "Attachment ID"
	// @Success 200
	// @Failure 400
	// @Failure 401
	// @Failure 404
	// @Router /media/attachments/{attachmentId} [get]
	mediaRouter.Get("/attachments/:attachmentId", MakeDownlodPostAttachmentHandler(mediaService))
}
