package service

import (
	"fmt"
	"go-read-var-log/config"
	"os"
	"regexp"
	"slices"
	"strings"
)

type GetLogResult struct {
	LogLines []string
	Strategy string
	Err      error
}

func errorGetLogResult(err error) GetLogResult {
	return GetLogResult{
		Err: err,
	}
}

func successGetLogResult(logLines []string, strategy string) GetLogResult {
	return GetLogResult{
		LogLines: logLines,
		Strategy: strategy,
	}
}

// GetLog returns the contents of a log file
// directoryPath: the path to the directory containing the log file
// filename: the name of the log file
// textMatch: a string to search for in the log file (case-sensitive, empty string if not required)
// regex: a compiled regular expression to search for in the log file (nil if not required)
// maxLines: the maximum number of lines to return (0 for all lines)
func GetLog(directoryPath string, filename string, textMatch string, regex *regexp.Regexp, maxLines int) GetLogResult {
	// TODO - REFACTOR: Consider a struct to hold the arguments to this function if the number of parameters grows

	filepath := strings.Join([]string{directoryPath, filename}, "/")

	if !validLogFromName(directoryPath, filename) {
		return errorGetLogResult(fmt.Errorf("invalid, unreadable or unsupported log file '%s'", filepath))
	}

	return selectLogStrategy(filepath)(filepath, textMatch, regex, maxLines)
}

func selectLogStrategy(filepath string) getLogStrategy {
	fileinfo, err := os.Stat(filepath)
	if err != nil {
		return func(filepath string, textMatch string, regex *regexp.Regexp, maxLines int) GetLogResult {
			return errorGetLogResult(fmt.Errorf("unable to stat file '%s': %s", filepath, err))
		}
	}
	filesize := fileinfo.Size()

	// TODO: Make this configurable
	if filesize < config.LargeFileBytes {
		return getSmallLog
	}

	// TODO: Add a strategy for large files
	return getSmallLog
}

type getLogStrategy func(filepath string, textMatch string, regex *regexp.Regexp, maxLines int) GetLogResult

func getSmallLog(filepath string, textMatch string, regex *regexp.Regexp, maxLines int) GetLogResult {
	// TODO: This is the simplest possible approach.  It will likely not work well for extremely large files.
	//       Consider seek() near the end of the file, backwards iteratively, until the desired number of lines is found.
	//       This will be more efficient for large files, but will be more complex to implement and maintain.
	//       On my machine (non-concurrent):
	//       - First scan of a 1GB file with 10.5 million lines takes ≈ 2-3s returning all (1) lines matching both
	//         a textMatch and regex.
	//       - Subsequent scans of the same file for a different textMatch and regex, returning all (1) matching lines,
	//         takes ≈ 1.5s.  This is likely due to the file fitting in the filesystem page cache on my system.
	//           kern.vm_page_free_min: 3500
	//           kern.vm_page_free_reserved: 912
	//           kern.vm_page_free_target: 4000
	byteSlice, err := os.ReadFile(filepath)
	if err != nil {
		return errorGetLogResult(err)
	}
	result := strings.Split(string(byteSlice), "\n")

	// Filter out lines that do not match the filters
	textMatchRequested := textMatch != ""
	regexMatchRequested := regex != nil
	if textMatchRequested || regexMatchRequested {
		// Single pass through the slice for efficiency
		result = slices.DeleteFunc(result, func(line string) bool {
			// Cheaper tests first, short circuit more expensive tests
			if textMatchRequested && !strings.Contains(line, textMatch) {
				return true
			}
			// Expensive tests last
			return regexMatchRequested && !regex.MatchString(line)
		})
	}

	// Restrict output to at most maxLines
	if maxLines > 0 && len(result) > maxLines {
		endIndex := len(result) - 1
		startIndex := max(0, endIndex-maxLines)
		result = result[startIndex:endIndex]
	}

	return successGetLogResult(result, "small")
}
