package posts

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func getPostId(c *fiber.Ctx) (int, error) {
	param := c.Params("postId")
	value, err := strconv.Atoi(param)
	if err != nil {
		log.Errorf("bad post id param: %s. err: %v", param, err)
		return 0, err
	}
	return value, nil
}
