package helper

import (
	"log"

	"git.martianoids.com/queru/retroserver/internal/sess"
)

type Color struct {
	Abbr   string
	Active bool
}

func (l *Color) Link() string {
	return "/pref/color/" + l.Abbr
}

// GetUserColor - Get user color
func (h *Helper) GetUserColor() string {
	s, err := sess.Get(h.Ctx)
	if err != nil {
		log.Println("GetUserColor: Cannot get sess")
		return "G"
	}

	if s.Get("color") != nil {
		return s.Get("color").(string)
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
	s, err := sess.Get(h.Ctx)
	if err != nil {
		log.Println("SetUserColor: Cannot get sess")
		return
	}
	defer s.Save()
	s.Set("color", color)
}

// IsActiveColor - bool true if color is this
func (h *Helper) IsActiveColor(color string) bool {
	return h.GetUserColor() == color
}
