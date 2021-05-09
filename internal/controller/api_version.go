package controller

import (
	"git.martianoids.com/queru/retroserver/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func APIVersionClient(c *fiber.Ctx) error {
	h := helper.New(c)
	h.FillDownloads()

	return c.JSON(fiber.Map{"latest": "v" + h.Downloads[0].Version})
}
