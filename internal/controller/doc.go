package controller

import "github.com/gofiber/fiber/v2"

func doc(app *fiber.App) {
	g := app.Group("/doc")

	g.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("doc/about", fiber.Map{
			"PageTitle": "About",
		}, "layouts/main")
	})

	g.Get("/contribute", func(c *fiber.Ctx) error {
		return c.Render("doc/contrib", fiber.Map{
			"PageTitle": "Contribute",
		}, "layouts/main")
	})
}
