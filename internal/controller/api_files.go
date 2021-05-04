package controller

import (
	"git.martianoids.com/queru/retroserver/internal/catalog"
	"github.com/gofiber/fiber/v2"
)

// controller
func apiFilesCtrl(app *fiber.App) {
	g := app.Group("/api/files")

	// file list
	g.Get("/catalog/:arch", func(c *fiber.Ctx) error {
		if !catalog.IsArch(c.Params("arch")) {
			return fiber.NewError(fiber.ErrNotFound.Code, "Architecture not found")
		}

		cat, err := catalog.New("./files/arch/"+c.Params("arch"), "", 0)
		if err != nil {
			return err
		}

		return c.JSON(cat)
	})

	// file download
	g.Get("/download/:arch/:id", func(c *fiber.Ctx) error {
		cat, err := catalog.New("./files/arch/"+c.Params("arch"), "", 0)
		if err != nil {
			return err
		}

		f, err := cat.Find(c.Params("id"))
		if err != nil {
			return fiber.NewError(fiber.ErrNotFound.Code, "File ID not found")
		}

		c.Set("Content-Disposition", "inline; filename=\""+f.Name+"\"")
		return c.SendFile(f.CompleteFileName("./files/arch/"+c.Params("arch")), true)
	})
}
