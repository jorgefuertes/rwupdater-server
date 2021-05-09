package helper

import "github.com/gofiber/fiber/v2"

// Render
func (h *Helper) Render(c *fiber.Ctx, tpl string) error {
	return c.Render(tpl, h, "layouts/main")
}
