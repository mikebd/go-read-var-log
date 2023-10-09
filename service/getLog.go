package service

import (
	"fmt"
	"os"
	"strings"
)

// GetLog returns the contents of a log file
// directoryPath: the path to the directory containing the log file
// filename: the name of the log file
// textMatch: a string to search for in the log file (case sensitive)
// maxLines: the maximum number of lines to return
func GetLog(directoryPath string, filename string, textMatch string, maxLines int) ([]string, error) {
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

	// Select lines that match the filters
	if textMatch != "" {
		filteredResult := make([]string, 0, len(result))
		for i, line := range result {
			if strings.Contains(line, textMatch) {
				filteredResult = append(filteredResult, result[i])
			}
		}
		result = filteredResult
	}

	// Restrict output to at most maxLines
	if maxLines > 0 {
		endIndex := len(result) - 1
		startIndex := max(0, endIndex-maxLines)
		result = result[startIndex:endIndex]
	}

	return result, nil
}
