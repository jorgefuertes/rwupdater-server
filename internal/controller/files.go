package controller

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const MaxRecursion = 10

// File
type File struct {
	ID        string `json:"id,omitempty"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Core      string `json:"core,omitempty"`
	Version   string `json:"version,omitempty"`
	Timestamp int64  `json:"ts"`
}

// aux function to transverse dirs
func recurse(arch string, path string, curRec int) ([]File, error) {
	var list []File
	var r = regexp.MustCompile(`\A([A-Za-z0-9\-]+)\_([0-9a-zA-Z\-\_]+)\.rbf\z`)
	var fbr = regexp.MustCompile(`.*\.fiber\...\z`)
	var prefix = "./files/" + arch

	if curRec > MaxRecursion {
		return list, errors.New("max recursion limit reached")
	}

	dir, err := ioutil.ReadDir(prefix + "/" + path)
	if err != nil {
		return list, err
	}

	for _, entry := range dir {
		if strings.HasPrefix(entry.Name(), ".") || fbr.MatchString(entry.Name()) {
			continue
		}

		if entry.IsDir() {
			subdir, err := recurse(arch, path+"/"+entry.Name(), curRec+1)
			if err != nil {
				return list, err
			}
			if len(subdir) > 0 {
				list = append(list, subdir...)
			}
			continue
		}

		file := File{
			Path:      strings.TrimPrefix(path, "/"),
			Name:      entry.Name(),
			Timestamp: entry.ModTime().Unix(),
		}

		digest := md5.New()
		digest.Write([]byte(file.Path + "/" + file.Name))
		file.ID = hex.EncodeToString(digest.Sum(nil))
		if r.MatchString(entry.Name()) {
			fdata := r.FindSubmatch([]byte(entry.Name()))
			file.Core = string(fdata[1])
			file.Version = string(fdata[2])
		}
		list = append(list, file)
	}

	return list, nil
}

// controller
func files(app *fiber.App) {
	f := app.Group("/files")

	// file list
	f.Get("/catalog/:arch", func(c *fiber.Ctx) error {
		dir, err := recurse(c.Params("arch"), "", 0)
		if err != nil {
			return err
		}

		return c.JSON(dir)
	})

	// file download
	f.Get("/download/:arch/:id", func(c *fiber.Ctx) error {
		dir, err := recurse(c.Params("arch"), "", 0)
		if err != nil {
			return err
		}

		var file File
		for _, file = range dir {
			if file.ID == c.Params("id") {
				break
			}
		}

		if file.ID != c.Params("id") {
			return fiber.ErrNotFound
		}

		c.Set("Content-Disposition", "inline; filename=\""+file.Name+"\"")
		return c.SendFile("./files/"+c.Params("arch")+"/"+file.Path+"/"+file.Name, true)
	})
}
