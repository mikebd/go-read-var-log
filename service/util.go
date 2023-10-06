package service

import (
	"go-read-var-log/config"
	"os"
	"strings"
)

// validLogFromDirectoryEntry returns true if the entry can be handled by this service
func validLogFromDirectoryEntry(directoryPath string, entry os.DirEntry) bool {
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

// validLogFromName returns true if the log can be handled by this service
func validLogFromName(directoryPath string, filename string) bool {
	filepath := strings.Join([]string{directoryPath, filename}, "/")

	// Cheaper tests first
	fileinfo, err := os.Stat(filepath)
	if fileinfo == nil || fileinfo.IsDir() || err != nil {
		return false
	}

	if !isFileSupported(filename) {
		return false
	}

	// More expensive tests last
	if !isFileReadable(filepath) {
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
