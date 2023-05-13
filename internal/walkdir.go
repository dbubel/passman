package internal

import (
	"os"
	"path/filepath"
)

func VisitFile(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			*files = append(*files, path)
		}
		return nil
	}
}

func GetAllFilesInDirectory() ([]string, error) {
	var files []string
	err := filepath.Walk(".", VisitFile(&files))
	if err != nil {
		return nil, err
	}
	return files, nil
}
