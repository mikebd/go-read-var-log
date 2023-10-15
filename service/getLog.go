package service

import (
	"fmt"
	"go-read-var-log/config"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

// GetLog returns the contents of a log file
func GetLog(params *GetLogParams) GetLogResult {
	const strategy = "unknown"
	if !validLogFromName(params.DirectoryPath, params.Filename) {
		return errorGetLogResult(strategy, fmt.Errorf("invalid, unreadable or unsupported log file '%s'", params.filepath()))
	}

	return selectLogStrategy(params.filepath())(params)
}

// selectLogStrategy returns the appropriate strategy for the log file
// TODO: Consider removing this function and getSmallLog() once getLargeLog() is implemented and thoroughly tested
func selectLogStrategy(filepath string) getLogStrategy {
	const strategy = "unknown"
	isFileLarge, err := isFileLarge(filepath)
	if err != nil {
		return func(params *GetLogParams) GetLogResult {
			return errorGetLogResult(strategy, fmt.Errorf("unable to stat file '%s': %s", filepath, err))
		}
	}

	if isFileLarge {
		return getLargeLog
	}

	return getSmallLog
}

// getLargeLog returns the contents of a large log file that does not easily fit in memory
func getLargeLog(params *GetLogParams) GetLogResult {
	const strategy = "large"

	// TODO: What if the result size is larger than the available memory?  e.g. no filters
	result := make([]string, 0, min(2000, params.MaxLines))

	// Open the file for reading
	file, errOpen := os.OpenFile(params.filepath(), os.O_RDONLY, 0)
	if errOpen != nil {
		return errorGetLogResult(strategy, fmt.Errorf("error opening file '%s': %s", params.filepath(), errOpen))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("error closing file '%s': %s\n", params.filepath(), err)
		}
	}(file)

	// Seek to the end of the file
	pos, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return errorGetLogResult(strategy, fmt.Errorf("error seeking to end of file '%s': %s", params.filepath(), err))
	}

	// Seek to the beginning of the next block to scan
	for pos > 0 {
		pos, err = file.Seek(-min(pos, config.SeekBytes), io.SeekCurrent)
		if err != nil {
			return errorGetLogResult(strategy, fmt.Errorf("error seeking to beginning of next block to scan in file '%s': %s", params.filepath(), err))
		}
		// log.Println("position:", pos)
	}

	// TODO: Need to handle partial lines at block boundaries

	return successGetLogResult(result, strategy)
}

// getSmallLog returns the contents of a small log file that easily fits in memory
func getSmallLog(params *GetLogParams) GetLogResult {
	const strategy = "small"
	byteSlice, err := os.ReadFile(params.filepath())
	if err != nil {
		return errorGetLogResult(strategy, err)
	}
	result := strings.Split(string(byteSlice), "\n")

	// Filter out lines that do not match the filters
	if params.matchRequested() {
		// Single pass through the slice for efficiency
		result = slices.DeleteFunc(result, func(line string) bool {
			// Cheaper tests first, short circuit more expensive tests
			if params.textMatchRequested() && !strings.Contains(line, params.TextMatch) {
				return true
			}
			// Expensive tests last
			return params.regexMatchRequested() && !params.Regex.MatchString(line)
		})
	}

	// Restrict output to at most maxLines
	if params.MaxLines > 0 && len(result) > params.MaxLines {
		endIndex := len(result) - 1
		startIndex := max(0, endIndex-params.MaxLines)
		result = result[startIndex:endIndex]
	}

	return successGetLogResult(result, strategy)
}
