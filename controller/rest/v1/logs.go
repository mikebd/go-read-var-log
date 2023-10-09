package v1

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-read-var-log/config"
	"go-read-var-log/controller/rest/util"
	"go-read-var-log/service"
	"log"
	"net/http"
	"slices"
	"strconv"
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
	defer log.Println(r.URL.RequestURI(), "took", time.Since(start))

	logFilename := p.ByName("log")

	maxLines := config.GetArguments().NumberOfLogLines
	maxLinesParam := r.URL.Query().Get("n")
	if maxLinesParam != "" {
		intParam, err := strconv.Atoi(maxLinesParam)
		if err == nil && intParam > 0 {
			maxLines = intParam
		} else {
			util.RenderTextPlain(w, nil, fmt.Errorf("invalid value for parameter 'n': '%s'", maxLinesParam))
			return
		}
	}

	logEvents, err := service.GetLog(config.LogDirectory, logFilename, maxLines)

	// Reverse the slice - we want the most recent events first.
	// Tests to see if this is slower than just iterating backwards when rendering
	// showed that it was not.  This is easier to read for maintainability.
	slices.Reverse(logEvents)

	util.RenderTextPlain(w, logEvents, err)
}
