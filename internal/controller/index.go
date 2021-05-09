package controller

import (
	"git.martianoids.com/queru/retroserver/internal/helper"
	"git.martianoids.com/queru/retroserver/internal/locale"
	"github.com/gofiber/fiber/v2"
)

func FrontIndex(c *fiber.Ctx) error {
	h := helper.New(c)
	h.SetPageTitle("menu.main.title")
	return h.Render(c, "index")
}

func Index(c *fiber.Ctx) error {
	return c.Redirect("/front/"+locale.GetUserLang(c)+"/amber", fiber.StatusTemporaryRedirect)
}
