package dao

import (
	"os"
	"path/filepath"
	"regexp"
)

func getFileNamesWithExt(ext string, where string) []string {
	var files []string
	filepath.Walk(where, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})
	return files
}
