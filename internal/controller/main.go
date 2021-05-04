package controller

import "github.com/gofiber/fiber/v2"

// Setup - Add all routes to app
func Setup(app *fiber.App) {
	indexCtrl(app)
	docCtrl(app)
	downloadCtrl(app)
	contactCtrl(app)
	langCtrl(app)
	apiServerCtrl(app)
	apiArchCtrl(app)
	apiFilesCtrl(app)
}
