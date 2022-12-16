package notes

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Save a Note to the disk with the title as filename
func SaveToDisk(n *Note, folder string) error {
	filename := filepath.Join(folder, n.Title)
	return os.WriteFile(filename, n.Body, 0600)
}

func searchKeyworidInFilename(folder string, keyword string) (string, error) {
	items, _ := os.ReadDir(folder)
	for _, item := range items {
		b, err := os.ReadFile(filepath.Join(folder, item.Name()))
		if err != nil {
			log.Printf("Could not read %v\n", item.Name())
		}
		s := string(b)

		if strings.Contains(s, keyword) {
			return item.Name(), nil
		}
	}

	return "", errors.New("no file contains this keyword")
}

// Scan files in a folder tofind first occurence of a keyword
func LoadFromDisk(keyword string, folder string) (*Note, error) {
	filename, err := searchKeyworidInFilename(folder, keyword)
	if err != nil {
		return nil, err
	}

	body, err := os.ReadFile(filepath.Join(folder, filename))
	if err != nil {
		return nil, err
	}

	return &Note{Title: filename, Body: body}, nil
}
