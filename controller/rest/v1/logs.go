package v1

import (
	"github.com/julienschmidt/httprouter"
	"go-read-var-log/config"
	"go-read-var-log/controller/rest/util"
	"go-read-var-log/service"
	"log"
	"net/http"
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
	defer log.Println(r.URL.RequestURI(), "took", time.Since(start))

	logFilename := p.ByName("log")

	maxLines, _ := util.PositiveIntParamStrict(w, r, config.GetArguments().NumberOfLogLines, "n")
	if maxLines > 0 {
		logEvents, err := service.GetLog(config.LogDirectory, logFilename, maxLines)

		if err == nil {
			// Reverse the slice - we want the most recent events first.
			// Tests to see if this is slower than just iterating backwards when rendering
			// showed that it was not.  This is easier to read for maintainability.
			slices.Reverse(logEvents)
		}

		util.RenderTextPlain(w, logEvents, err)
	}
}
