package media

import (
	"bytes"
	"errors"
	"io"
	"strconv"

	"github.com/PaulShpilsher/instalike/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

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
func MakeDownlodPostAttachmentHandler(s MediaService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		attachmentId, err := getAttachmentId(c)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		// TODO: in future rewrite that the attachments are stored in a separate storage
		// and database contants only attachment's metadata
		attachment, err := s.GetPostAttachment(attachmentId)
		if err != nil {
			if errors.Is(err, utils.ErrNotFound) {
				return c.SendStatus(fiber.StatusNotFound)
			}

			log.Error(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// we're gonna to pretend that the stream  is comming from some other storage
		// but right now lets simulate that with creating byte reader from byte slice
		reader := bytes.NewReader(attachment.Data)
		// byte reader doesn't have Close(), so the next line is a placeholder when we have a file reader
		// defer reader.Close()

		// TODO: implement streaming.  for now send the whole thing
		c.Set(fiber.HeaderContentType, attachment.ContentType)
		c.Set(fiber.HeaderContentLength, strconv.FormatInt(int64(attachment.Size), 10))
		_, err = io.Copy(c.Response().BodyWriter(), reader)
		if err != nil {
			log.Errorf("copying to response failed", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func getAttachmentId(c *fiber.Ctx) (int, error) {
	param := c.Params("attachmentId")
	value, err := strconv.Atoi(param)
	if err != nil {
		log.Errorf("bad attachment id param: %s. err: %v", param, err)
		return 0, err
	}
	return value, nil
}
