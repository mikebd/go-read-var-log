package rest

import (
	"github.com/julienschmidt/httprouter"
	"go-read-var-log/config"
	"go-read-var-log/controller/rest/v1"
	"net/http"
	"strconv"
)

func StartHttpRouter(args *config.Arguments) error {
	router := httprouter.New()
	router.GET("/api/v1/logs", v1.GetLogs)
	router.GET("/api/v1/logs/:log", v1.GetLog)
	return http.ListenAndServe(":"+strconv.Itoa(args.HttpPort), router)
}
