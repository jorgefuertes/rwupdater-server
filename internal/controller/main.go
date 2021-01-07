package controller

import "github.com/gofiber/fiber/v2"

// Setup - Add all routes to app
func Setup(app *fiber.App) {
	index(app)
	server(app)
	files(app)
}
