package controller

import (
	"git.martianoids.com/queru/retroserver/internal/catalog"
	"github.com/gofiber/fiber/v2"
)

// APIArch - Arch list
func APIArch(c *fiber.Ctx) error {
	list, err := catalog.ArchList()
	if err != nil {
		return err
	}

	return c.JSON(list)
}
