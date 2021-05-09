package controller

import (
	"fmt"

	"git.martianoids.com/queru/retroserver/internal/helper"
	"github.com/gofiber/fiber/v2"
)

// Download
type Download struct {
	File    string
	Version string
	Os      string
	CPU     string
	Ext     string
	Size    string
}

func FrontDownloads(c *fiber.Ctx) error {
	h := helper.New(c)
	h.SetPageTitle("menu.downloads.title")
	h.FillDownloads()
	return h.Render(c, "downloads")
}

func DownloadHandler(c *fiber.Ctx) error {
	h := helper.New(c)
	if err := h.FillDownloads(); err != nil {
		return err
	}
	d, err := h.FindDownload(c.Params("os"), c.Params("arch"))
	if err != nil {
		return err
	}

	if c.Params("mode") == "dist" {
		c.Set("Content-Disposition", "inline; filename=\""+d.Dist.File+"\"")
		c.Set("Content-Type", "application/zip")
		return c.SendFile("./files/client/dist/"+d.Dist.File, false)
	}

	if c.Params("mode") == "bin" {
		c.Set("Content-Disposition", "inline; filename=\""+d.Bin.File+"\"")
		c.Set("Content-Type", "application/x-executable")
		return c.SendFile(
			fmt.Sprintf("./files/client/bin/%s/%s/%s", d.Os, d.CPU, d.Bin.File),
			false,
		)
	}

	return fiber.NewError(fiber.ErrNotFound.Code, "File not found")
}
