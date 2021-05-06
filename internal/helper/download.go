package helper

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/gofiber/fiber/v2"
)

// Download
type Download struct {
	Dist struct {
		File string
		Ext  string
		Size string
	}
	Bin struct {
		File string
		Size string
	}
	Version string
	Os      string
	CPU     string
}

// DistLink - Download path
func (d *Download) DistLink() string {
	return fmt.Sprintf("/download/dist/%s/%s", d.Os, d.CPU)
}

// BinLink - Download path
func (d *Download) BinLink() string {
	return fmt.Sprintf("/download/bin/%s/%s", d.Os, d.CPU)
}

// NameDistLink - Download name
func (d *Download) DistName() string {
	return fmt.Sprintf("%s %s", d.Dist.Ext, d.Dist.Size)
}

// NameBinLink - Download name
func (d *Download) BinName() string {
	return d.Bin.Size
}

// FillDownloads
func (h *Helper) FillDownloads() error {
	h.Downloads = make([]Download, 0)
	dir, err := os.ReadDir("./files/client/dist")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "error reading dist dir")
	}

	r := regexp.MustCompile(`_v([0-9\.]+)-([a-z]+)_([a-z0-9]+)\.(gz|bz|zip)`)
	for _, f := range dir {
		if !r.Match([]byte(f.Name())) {
			continue
		}
		d := new(Download)
		// dist
		d.Dist.File = f.Name()
		fdata := r.FindSubmatch([]byte(d.Dist.File))
		d.Version = string(fdata[1])
		d.Os = string(fdata[2])
		d.CPU = string(fdata[3])
		d.Dist.Ext = strings.ToTitle(string(fdata[4]))
		fi, _ := f.Info()
		d.Dist.Size = humanize.Bytes(uint64(fi.Size()))
		// bin
		d.Bin.File = "rw-updater"
		if d.Os == "windows" {
			d.Bin.File += ".exe"
		}
		fi, err = os.Stat(
			fmt.Sprintf("./files/client/bin/%s/%s/%s", d.Os, d.CPU, d.Bin.File))
		if err == nil {
			d.Bin.Size = humanize.Bytes(uint64(fi.Size()))
		}
		h.Downloads = append(h.Downloads, *d)
	}
	h.Latest = h.Downloads[0].Version

	return nil
}

// FindDownload
func (h *Helper) FindDownload(o, a string) (Download, error) {
	for _, d := range h.Downloads {
		if d.Os == o && d.CPU == a {
			return d, nil
		}
	}

	return Download{}, fiber.NewError(fiber.ErrNotFound.Code, "Download not found")
}
