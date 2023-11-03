package v1

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-read-var-log/config"
	"go-read-var-log/controller/rest/util"
	"go-read-var-log/service"
	"log"
	"net/http"
	"regexp"
	"slices"
	"time"
)

// GetLogs handles GET /api/v1/logs and returns a list of log files in the log directory
func GetLogs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	start := time.Now()
	defer log.Println(r.URL.RequestURI(), "took", time.Since(start))

	logFilenames, err := service.ListLogs(config.LogDirectory)
	util.RenderTextPlain(w, logFilenames, err)
}

// GetLog handles GET /api/v1/logs/{log} and return the contents in reverse order
func GetLog(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	start := time.Now()
	linesReturned := 0
	strategy := "unknown"
	defer func() {
		log.Println(r.URL.RequestURI(), "returned", linesReturned, "lines using", strategy, "strategy in", time.Since(start))
	}()

	logFilename := p.ByName("log")

	textMatch := r.URL.Query().Get("q")
	regexPattern := r.URL.Query().Get("r")
	var regex *regexp.Regexp
	if regexPattern != "" {
		regexCompiled, err := regexp.Compile(regexPattern)
		if err != nil {
			util.RenderTextPlain(w, nil, fmt.Errorf("invalid regex pattern '%s': %s", regexPattern, err))
			return
		}
		regex = regexCompiled
	}

	maxLines, err := util.PositiveIntParamStrict(w, r, config.GetArguments().NumberOfLogLines, "n")
	if err == nil {
		getLogResult := service.GetLog(&service.GetLogParams{
			DirectoryPath: config.LogDirectory,
			Filename:      logFilename,
			TextMatch:     textMatch,
			Regex:         regex,
			MaxLines:      maxLines,
		})
		logEvents, err := getLogResult.LogLines, getLogResult.Err
		strategy = getLogResult.Strategy

		if err == nil {
			linesReturned = len(logEvents)

			// Reverse the slice - we want the most recent events first.
			// Tests to see if this is slower than just iterating backwards when rendering
			// showed that it was not.  This is easier to read for maintainability.
			slices.Reverse(logEvents)
		}

		util.RenderTextPlain(w, logEvents, err)
	}
}
