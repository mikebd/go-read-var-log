package service

import (
	"os"
)

// ListLogs returns a list of log files in the log directory that are readable by this service
func ListLogs(directoryPath string) ([]string, error) {
	result := []string{}

	entries, err := os.ReadDir(directoryPath)
	if err != nil {
		return nil, err
	}

	/* filter out entries that are not readable or supported */
	for _, entry := range entries {
		if validLogForThisService(directoryPath, entry) {
			println("Adding entry:", entry.Name())
			result = append(result, entry.Name())
		}
	}

	return result, nil
}
