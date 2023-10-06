package v1

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-read-var-log/config"
	"go-read-var-log/service"
	"log"
	"net/http"
	"strings"
)

func GetLogs(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	log.Println("handling /api/v1/logs")

	logFilenames, err := service.ListLogs(config.LogDirectory)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if err == nil {
		_, err = fmt.Fprintln(w, strings.Join(logFilenames, "\n"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = fmt.Fprintln(w, err.Error())
	}
}
