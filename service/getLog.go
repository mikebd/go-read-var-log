package service

import (
	"fmt"
	"go-read-var-log/config"
	"os"
	"slices"
	"strings"
)

// GetLog returns the contents of a log file
func GetLog(params *GetLogParams) GetLogResult {
	if !validLogFromName(params.DirectoryPath, params.Filename) {
		return errorGetLogResult(fmt.Errorf("invalid, unreadable or unsupported log file '%s'", params.filepath()))
	}

	return selectLogStrategy(params.filepath())(params)
}

func selectLogStrategy(filepath string) getLogStrategy {
	fileinfo, err := os.Stat(filepath)
	if err != nil {
		return func(params *GetLogParams) GetLogResult {
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

func getSmallLog(params *GetLogParams) GetLogResult {
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
	byteSlice, err := os.ReadFile(params.filepath())
	if err != nil {
		return errorGetLogResult(err)
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

	return successGetLogResult(result, "small")
}
