package controller

import (
	"github.com/gofiber/fiber/v2"
)

func FrontError(c *fiber.Ctx) error {
	if c.Params("code") == "404" {
		return fiber.NewError(fiber.ErrNotFound.Code, "Resource not found")
	}
	return fiber.NewError(fiber.ErrInternalServerError.Code, "Internal server error")
}
