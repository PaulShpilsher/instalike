package media

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(mediaRouter fiber.Router, authMiddleware fiber.Handler, mediaService MediaService) {

	// DownloadPostMultimediAttachment godoc
	// @Summary Downloads multimedia file by post's attachmentId
	// @Description downloads multimedia file by post's attachmentId
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
	mediaRouter.Get("/media/attachments/:attachmentId", authMiddleware, MakeDownlodPostAttachmentHandler(mediaService))

}
