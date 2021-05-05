package helper

import (
	"html/template"

	"git.martianoids.com/queru/retroserver/internal/cfg"
	"git.martianoids.com/queru/retroserver/internal/locale"
	"github.com/gofiber/fiber/v2"
)

type Color struct {
	Abbr   string
	Active bool
}

func (l *Color) Link() string {
	return "/pref/color/" + l.Abbr
}

type Lang struct {
	Abbr   string
	Active bool
}

func (l *Lang) Link() string {
	return "/pref/lang/" + l.Abbr
}

// Download
type Download struct {
	File    string
	Version string
	Os      string
	CPU     string
	Ext     string
	Size    string
}

// Link to download
func (d *Download) Link() string {
	return "/download/dist/" + d.File
}

// Helper
type Helper struct {
	Ctx       *fiber.Ctx
	Lang      string
	PageTitle template.HTML
	Vars      fiber.Map
	IsMain    bool
	Downloads []Download
	Latest    string
	Err       string
	Colors    []Color
	Langs     []Lang
}

// New - Create and return new Helper
func New(c *fiber.Ctx) Helper {
	h := new(Helper)
	h.Ctx = c
	h.Colors = make([]Color, 0)
	h.Colors = append(h.Colors,
		Color{"W", h.IsActiveColor("W")},
		Color{"G", h.IsActiveColor("G")},
		Color{"B", h.IsActiveColor("B")},
		Color{"A", h.IsActiveColor("A")},
	)
	h.Lang = locale.GetUserLang(c)
	h.Langs = make([]Lang, 0)
	h.Langs = append(h.Langs, Lang{"en", h.IsActiveLang("en")}, Lang{"es", h.IsActiveLang("es")})
	h.Vars = make(fiber.Map)
	return *h
}

// SetPageTile - Set the current page title
func (h *Helper) SetPageTitle(key string) {
	h.PageTitle = h.T(key)
	h.IsMain = (key == "menu.main.title")
}

// T - Translate
func (h *Helper) T(key string, arg ...interface{}) template.HTML {
	return locale.T(h.Lang, key, arg...)
}

// IsLang - True if Lang == param
func (h *Helper) IsActiveLang(l string) bool {
	return (h.Lang == l)
}

// GetUserColor - Get user color
func (h *Helper) GetUserColor() string {
	sess, err := cfg.Session.Get(h.Ctx)
	if err != nil {
		panic(err)
	}

	if sess.Get("color") != nil {
		return sess.Get("color").(string)
	}

	// default to green
	return "G"
}

// GetColorCSS
func (h *Helper) LinkColorCSS() string {
	return "/asset/css/color/" + h.GetUserColor() + ".css"
}

// SetColor - Set user color
func (h *Helper) SetUserColor(color string) {
	sess, err := cfg.Session.Get(h.Ctx)
	if err != nil {
		panic(err)
	}
	defer sess.Save()
	sess.Set("color", color)
}

// IsActiveColor - bool true if color is this
func (h *Helper) IsActiveColor(color string) bool {
	return h.GetUserColor() == color
}
