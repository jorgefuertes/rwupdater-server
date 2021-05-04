package controller

import (
	"os"
	"regexp"
	"strings"

	"git.martianoids.com/queru/retroserver/internal/helper"
	"github.com/dustin/go-humanize"
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

func downloadCtrl(app *fiber.App) {
	app.Get("/downloads", func(c *fiber.Ctx) error {
		h := helper.New(c)
		h.SetPageTitle("menu.downloads.title")

		r := regexp.MustCompile(`_v([0-9\.]+)-([a-z]+)_([a-z0-9]+)\.(gz|bz|zip)`)

		h.Downloads = make([]helper.Download, 0)
		dir, err := os.ReadDir("./files/client/dist")
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "error reading dist dir")
		}
		for _, f := range dir {
			if !r.Match([]byte(f.Name())) {
				continue
			}
			d := new(helper.Download)
			d.File = f.Name()
			fdata := r.FindSubmatch([]byte(d.File))
			d.Version = string(fdata[1])
			d.Os = string(fdata[2])
			d.CPU = string(fdata[3])
			d.Ext = strings.ToTitle(string(fdata[4]))
			fi, _ := f.Info()
			d.Size = humanize.Bytes(uint64(fi.Size()))
			h.Downloads = append(h.Downloads, *d)
		}
		h.Latest = h.Downloads[0].Version

		return h.Render("downloads")
	})

	app.Get("/download/dist/:fname", func(c *fiber.Ctx) error {
		if _, err := os.Stat("./files/client/dist/" + c.Params("fname")); err != nil {
			return fiber.ErrNotFound
		}

		c.Set("Content-Disposition", "inline; filename=\""+c.Params("fname")+"\"")
		return c.SendFile("./files/client/dist/"+c.Params("fname"), false)
	})
}
