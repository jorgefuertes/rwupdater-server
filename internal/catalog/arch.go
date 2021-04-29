package catalog

import (
	"os"
)

// ArchList - Architectures list
func ArchList() ([]string, error) {
	list := make([]string, 0)

	dir, err := os.ReadDir("./files/arch")
	if err != nil {
		return list, err
	}

	for _, f := range dir {
		if f.IsDir() {
			list = append(list, f.Name())
		}
	}

	return list, nil
}

// IsArch - Check if arch exists
func IsArch(arch string) bool {
	list, _ := ArchList()
	for _, a := range list {
		if a == arch {
			return true
		}
	}

	return false
}
