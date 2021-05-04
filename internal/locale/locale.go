package locale

import (
	"html/template"
	"strings"

	"git.martianoids.com/queru/retroserver/internal/cfg"
	"github.com/gofiber/fiber/v2"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/yaml"
)

// I18n - Translations
var I18n *i18n.I18n

func load() {
	I18n = i18n.New(yaml.New("locales"))
}

func init() {
	load()
}

// Refresh - Reload yaml locales
func Refresh() {
	load()
}

// SetUserLang - Set the user's session language.
func SetUserLang(c *fiber.Ctx, l string) {
	sess, err := cfg.Session.Get(c)
	if err != nil {
		panic(err)
	}
	defer sess.Save()
	sess.Set("lang", l)
}

// GetUserLang - Returns the user's session lang. English as default.
func GetUserLang(c *fiber.Ctx) string {
	sess, err := cfg.Session.Get(c)
	if err != nil {
		panic(err)
	}

	if sess.Get("lang") != nil {
		return sess.Get("lang").(string)
	} else {
		if strings.HasPrefix(c.Get("User-Agent"), "es") {
			return "es"
		}
	}

	return "en"
}

// T - Translate
func T(lang string, key string, args ...interface{}) template.HTML {
	return I18n.Default("["+lang+"] "+key+"!").T(lang, key, args...)
}
