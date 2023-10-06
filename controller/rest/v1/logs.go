package v1

import (
	"github.com/julienschmidt/httprouter"
	"go-read-var-log/config"
	"go-read-var-log/controller/rest/util"
	"go-read-var-log/service"
	"log"
	"net/http"
)

// GetLogs handles GET /api/v1/logs and returns a list of log files in the log directory
func GetLogs(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	log.Println("handling /api/v1/logs")
	logFilenames, err := service.ListLogs(config.LogDirectory)
	util.RenderTextPlain(w, logFilenames, err)
}
