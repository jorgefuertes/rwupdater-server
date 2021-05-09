package helper

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Color struct {
	Name   string
	Abbr   string
	Link   string
	Active bool
}

func (h *Helper) fillColors(c *fiber.Ctx) {
	h.Color = c.Params("color")
	h.Colors = make([]Color, 0)
	if h.Color == "" {
		h.Color = "amber"
	}
	for _, name := range colors {
		h.Colors = append(h.Colors,
			Color{
				Name:   name,
				Abbr:   strings.ToUpper(string(name[0])),
				Link:   h.colorLink(c, name),
				Active: (h.Color == name),
			})
	}
}

func (h *Helper) colorLink(c *fiber.Ctx, name string) string {
	return "/front/" + c.Params("lang") + "/" + name
}

// GetColorCSS
func (h *Helper) LinkColorCSS() string {
	return "/asset/css/color/" + h.Color + ".css"
}
