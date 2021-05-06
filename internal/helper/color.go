package helper

import "git.martianoids.com/queru/retroserver/internal/cfg"

type Color struct {
	Abbr   string
	Active bool
}

func (l *Color) Link() string {
	return "/pref/color/" + l.Abbr
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
