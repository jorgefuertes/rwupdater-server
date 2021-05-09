package controller

import (
	"git.martianoids.com/queru/retroserver/internal/helper"
	"github.com/gofiber/fiber/v2"
)

// document controller
func FrontDoc(c *fiber.Ctx) error {
	h := helper.New(c)
	h.SetPageTitle("menu." + c.Params("doc") + ".title")
	return h.Render(c, "doc/"+c.Params("doc"))
}
