package locale

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"strings"

	"git.martianoids.com/queru/retroserver/internal/sess"
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

// GetUserLang - Returns the user's session lang. English as default.
func GetUserLang(c *fiber.Ctx) string {
	s, err := sess.Get(c)
	if err != nil {
		log.Println("GetUserLang: Cannot get sess")
		return "es"
	}

	if s.Get("lang") != nil {
		return s.Get("lang").(string)
	} else {
		if strings.HasPrefix(c.Get("Accept-Language"), "es") {
			return "es"
		}
	}

	return "en"
}

// T - Translate
func T(lang string, key string, args ...interface{}) template.HTML {
	return I18n.Default("["+lang+"] "+key+"!").T(lang, key, args...)
}
