package service

import (
	"go-read-var-log/config"
	"os"
	"strings"
)

// isFileLarge returns true if the file is larger than the configured threshold
func isFileLarge(filepath string) (bool, error) {
	fileinfo, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}
	filesize := fileinfo.Size()

	largeFileBytes := func() int64 {
		if config.GetArguments() == nil {
			return config.LargeFileBytes
		}
		return config.GetArguments().LargeFileBytes
	}()

	return filesize >= largeFileBytes, nil
}

// isFileReadable returns true if the file is readable
func isFileReadable(filepath string) bool {
	_, err := os.Open(filepath)
	return err == nil
}

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

// isFileSupported returns true if the file type is supported by this service
func isFileSupported(filename string) bool {
	extension := strings.ToLower(filename[strings.LastIndex(filename, ".")+1:])
	return !strings.Contains(config.UnsupportedFileTypes, extension)
}
