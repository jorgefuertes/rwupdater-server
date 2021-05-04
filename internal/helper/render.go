package helper

// Render
func (h *Helper) Render(tpl string) error {
	return h.Ctx.Render(tpl, h, "layouts/main")
}
