package helper

import (
	"fmt"
	"html/template"

	"github.com/gofiber/fiber/v2"
)

var colors = []string{"white", "green", "blue", "amber"}

// Helper
type Helper struct {
	PageTitle template.HTML
	Vars      fiber.Map
	IsMain    bool
	Downloads []Download
	Latest    string
	Err       string
	Colors    []Color
	Color     string
	Langs     []Lang
	Lang      string
	MyURL     string
}

// New - Create and return new Helper
func New(c *fiber.Ctx) Helper {
	h := new(Helper)
	h.fillLangs(c)
	h.fillColors(c)
	// url
	h.MyURL = c.BaseURL()
	// free controller vars
	h.Vars = make(fiber.Map)
	return *h
}

func (h *Helper) LinkTo(p string) string {
	return fmt.Sprintf("/front/%s/%s/%s", h.Lang, h.Color, p)
}
