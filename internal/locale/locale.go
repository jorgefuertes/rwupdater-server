package locale

import (
	"embed"
	"html/template"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/yaml"
)

//go:embed locales
var locales embed.FS

// I18n - Translations
var I18n *i18n.I18n

func load() {
	I18n = i18n.New(yaml.NewWithFilesystem(http.FS(locales)))
}

func init() {
	load()
}

// Refresh - Reload yaml locales
func Refresh() {
	load()
}

// GetUserLang - Returns the user's browser lang.
func GetUserLang(c *fiber.Ctx) string {
	if c.Params("lang") != "" {
		return c.Params("lang")
	}
	if strings.HasPrefix(c.Get("Accept-Language"), "en") {
		return "en"
	}

	return "es"
}

// T - Translate
func T(lang string, key string, args ...interface{}) template.HTML {
	return I18n.Default("["+lang+"] "+key+"!").T(lang, key, args...)
}
