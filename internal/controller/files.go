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

func recurse(path string, curRec int) ([]File, error) {
	var list []File
	var r = regexp.MustCompile(`\A([A-Za-z0-9\-]+)\_([0-9a-z]+)\.rbf\z`)

	if curRec > MaxRecursion {
		return list, errors.New("max recursion limit reached")
	}

	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return list, err
	}

	for _, entry := range dir {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		if entry.IsDir() {
			subdir, err := recurse(path+"/"+entry.Name(), curRec+1)
			if err != nil {
				return list, err
			}
			if len(subdir) > 0 {
				list = append(list, subdir...)
			}
			continue
		}

		file := File{
			Path:      strings.TrimPrefix(path, "./files"),
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

func files(app *fiber.App) {
	f := app.Group("/files")

	// file list
	f.Get("/catalog/:arch", func(c *fiber.Ctx) error {
		dir, err := recurse("./files/"+c.Params("arch"), 0)
		if err != nil {
			return err
		}

		return c.JSON(dir)
	})

	// file download
	f.Get("/download/:id", func(c *fiber.Ctx) error {
		dir, err := recurse("./files", 0)
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
		return c.SendFile(file.Path+"/"+file.Name, true)
	})

	// core
	f.Get("/core/:name/:version", func(c *fiber.Ctx) error {
		dir, err := ioutil.ReadDir("./files/cores")
		if err != nil {
			return err
		}

		r := regexp.MustCompile(`\A(` + c.Params("name") + `)\_([0-9a-z]+)\.rbf\z`)
		for _, file := range dir {
			if r.MatchString(file.Name()) {
				fdata := r.FindSubmatch([]byte(file.Name()))
				if string(fdata[2]) == c.Params("version") {
					return c.SendStatus(fiber.StatusNotModified)
				}

				c.Set("Content-Disposition", "inline; filename=\""+file.Name()+"\"")
				return c.SendFile("./files/cores/"+file.Name(), true)
			}
		}

		return fiber.ErrNotFound
	})
}
