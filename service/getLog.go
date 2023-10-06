package service

import (
	"fmt"
	"os"
	"strings"
)

// GetLog returns the contents of a log file
func GetLog(directoryPath string, filename string) ([]string, error) {
	filepath := strings.Join([]string{directoryPath, filename}, "/")

	if !validLogFromName(directoryPath, filename) {
		return nil, fmt.Errorf("invalid, unreadable or unsupported log file '%s'", filepath)
	}

	// TODO: this is the simplest possible approach.  It will likely not work well for large files.
	byteSlice, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(byteSlice), "\n"), nil
}
