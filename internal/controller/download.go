package controller

import (
	"os"
	"regexp"
	"strings"

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

func download(app *fiber.App) {
	app.Get("/downloads", func(c *fiber.Ctx) error {
		r := regexp.MustCompile(`_v([0-9\.]+)-([a-z]+)_([a-z0-9]+)\.(gz|bz|zip)`)
		dls := make([]Download, 0)
		dir, err := os.ReadDir("./files/client/dist")
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "error reading dist dir")
		}
		for _, f := range dir {
			if !r.Match([]byte(f.Name())) {
				continue
			}
			d := new(Download)
			d.File = f.Name()
			fdata := r.FindSubmatch([]byte(d.File))
			d.Version = string(fdata[1])
			d.Os = string(fdata[2])
			d.CPU = string(fdata[3])
			d.Ext = strings.ToTitle(string(fdata[4]))
			fi, _ := f.Info()
			d.Size = humanize.Bytes(uint64(fi.Size()))
			dls = append(dls, *d)
		}

		return c.Render("downloads", fiber.Map{
			"PageTitle": "Downloads",
			"Latest":    dls[0].Version,
			"Dls":       dls,
		}, "layouts/main")
	})

	app.Get("/download/dist/:fname", func(c *fiber.Ctx) error {
		if _, err := os.Stat("./files/client/dist/" + c.Params("fname")); err != nil {
			return fiber.ErrNotFound
		}

		c.Set("Content-Disposition", "inline; filename=\""+c.Params("fname")+"\"")
		return c.SendFile("./files/client/dist/"+c.Params("fname"), false)
	})
}
