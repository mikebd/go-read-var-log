package service

import (
	"os"
)

// ListLogs returns a list of log files in the log directory that are readable by this service
func ListLogs(directoryPath string) ([]string, error) {
	var result []string

	entries, err := os.ReadDir(directoryPath)
	if err != nil {
		return nil, err
	}

	/* filter out entries that are not readable or supported */
	for _, entry := range entries {
		if validLogForThisService(directoryPath, entry) {
			result = append(result, entry.Name())
		}
	}

	return result, nil
}
