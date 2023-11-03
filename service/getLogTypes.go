package service

import (
	"regexp"
	"strings"
)

// GetLogParams provides the parameters for GetLog()
// DirectoryPath: the path to the directory containing the log file
// Filename: the name of the log file
// TextMatch: a string to search for in the log file (case-sensitive, empty string if not required)
// Regex: a compiled regular expression to search for in the log file (nil if not required)
// MaxLines: the maximum number of lines to return (0 for all lines)
type GetLogParams struct {
	DirectoryPath string
	Filename      string
	TextMatch     string
	Regex         *regexp.Regexp
	MaxLines      int
}

func (p *GetLogParams) filepath() string {
	return strings.Join([]string{p.DirectoryPath, p.Filename}, "/")
}

func (p *GetLogParams) textMatchRequested() bool {
	return p.TextMatch != ""
}

func (p *GetLogParams) regexMatchRequested() bool {
	return p.Regex != nil
}

type GetLogResult struct {
	LogLines []string
	Strategy string
	Err      error
}

func errorGetLogResult(strategy string, err error) GetLogResult {
	return GetLogResult{
		Strategy: strategy,
		Err:      err,
	}
}

func successGetLogResult(logLines []string, strategy string) GetLogResult {
	return GetLogResult{
		LogLines: logLines,
		Strategy: strategy,
	}
}

type getLogStrategy func(params *GetLogParams) GetLogResult
