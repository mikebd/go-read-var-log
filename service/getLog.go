package service

import (
	"fmt"
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
	return errorGetLogResult(strategy, fmt.Errorf("large files are not yet supported"))
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
