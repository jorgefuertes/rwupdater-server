package catalog

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"os"
	"regexp"
	"strings"
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

// Catalog
type Catalog []File

// CompletePath
func (f *File) CompletePath(root string) string {
	return root + "/" + f.Path
}

// CompleteFileName
func (f *File) CompleteFileName(root string) string {
	return root + "/" + f.Path + "/" + f.Name
}

// Find - Find by id
func (c *Catalog) Find(id string) (File, error) {
	for _, f := range *c {
		if f.ID == id {
			return f, nil
		}
	}

	return File{}, errors.New("file not found")
}

// PathExists - Bool
func (c *Catalog) PathExists(path string) bool {
	for _, f := range *c {
		if f.Path == path {
			return true
		}
	}

	return false
}

// FindByName - Find one by path and name
func (c *Catalog) FindByName(path, name string) (File, error) {
	for _, f := range *c {
		if f.Path == path && f.Name == name {
			return f, nil
		}
	}

	return File{}, errors.New("file not found")
}

// New
func New(root string, path string, curRec int) (*Catalog, error) {
	var cat Catalog
	var r = regexp.MustCompile(`\A([A-Za-z0-9\-]+)\_([0-9a-zA-Z\-\_]+)\.(rbf|ua2|np1)\z`)
	var fbr = regexp.MustCompile(`.*\.fiber\...\z`)

	if curRec > MaxRecursion {
		return &cat, errors.New("max recursion limit reached")
	}

	dir, err := os.ReadDir(root + "/" + path)
	if err != nil {
		return &cat, err
	}

	for _, entry := range dir {
		if strings.HasPrefix(entry.Name(), ".") || fbr.MatchString(entry.Name()) {
			continue
		}

		if entry.IsDir() {
			subdir, err := New(root, path+"/"+entry.Name(), curRec+1)
			if err != nil {
				return &cat, err
			}
			if len(*subdir) > 0 {
				cat = append(cat, *subdir...)
			}
			continue
		}

		info, _ := entry.Info()
		file := File{
			Path:      strings.TrimPrefix(path, "/"),
			Name:      entry.Name(),
			Timestamp: info.ModTime().Unix(),
		}

		digest := md5.New()
		digest.Write([]byte(file.Path + "/" + file.Name))
		file.ID = hex.EncodeToString(digest.Sum(nil))
		if r.MatchString(entry.Name()) {
			fdata := r.FindSubmatch([]byte(entry.Name()))
			file.Core = string(fdata[1])
			file.Version = string(fdata[2])
		}
		cat = append(cat, file)
	}

	return &cat, nil
}
