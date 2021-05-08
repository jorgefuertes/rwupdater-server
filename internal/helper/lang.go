package helper

import (
	"html/template"

	"git.martianoids.com/queru/retroserver/internal/locale"
)

type Lang struct {
	Abbr   string
	Active bool
}

func (l *Lang) Link() string {
	return "?lang=" + l.Abbr
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
