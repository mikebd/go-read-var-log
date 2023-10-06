package service

import (
	"go-read-var-log/config"
	"os"
	"strings"
)

func validLogForThisService(directoryPath string, entry os.DirEntry) bool {
	// Cheaper tests first
	if entry.IsDir() {
		return false
	}

	if !isFileSupported(entry.Name()) {
		return false
	}

	// More expensive tests last
	if !isFileReadable(strings.Join([]string{directoryPath, entry.Name()}, "/")) {
		return false
	}

	return true
}

func isFileReadable(path string) bool {
	_, err := os.Open(path)
	return err == nil
}

func isFileSupported(filename string) bool {
	extension := strings.ToLower(filename[strings.LastIndex(filename, ".")+1:])
	return !strings.Contains(config.UnsupportedFileTypes, extension)
}
