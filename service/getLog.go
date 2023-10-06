package service

import (
	"fmt"
	"os"
	"strings"
)

// GetLog returns the contents of a log file
func GetLog(directoryPath string, filename string, maxLines int) ([]string, error) {
	filepath := strings.Join([]string{directoryPath, filename}, "/")

	if !validLogFromName(directoryPath, filename) {
		return nil, fmt.Errorf("invalid, unreadable or unsupported log file '%s'", filepath)
	}

	// TODO: this is the simplest possible approach.  It will likely not work well for large files.
	byteSlice, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	result := strings.Split(string(byteSlice), "\n")

	// Restrict output to at most maxLines
	endIndex := len(result) - 1
	startIndex := max(0, endIndex-maxLines)
	result = result[startIndex:endIndex]

	return result, nil
}
