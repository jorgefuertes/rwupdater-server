package controller

import (
	"git.martianoids.com/queru/retroserver/internal/catalog"
	"github.com/gofiber/fiber/v2"
)

// APIFileList
func APIFileList(c *fiber.Ctx) error {
	if !catalog.IsArch(c.Params("arch")) {
		return fiber.NewError(fiber.ErrNotFound.Code, "Architecture not found")
	}

	cat, err := catalog.New("./files/arch/"+c.Params("arch"), "", 0)
	if err != nil {
		return err
	}

	return c.JSON(cat)
}

// APIDownload
func APIDownload(c *fiber.Ctx) error {
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
}
