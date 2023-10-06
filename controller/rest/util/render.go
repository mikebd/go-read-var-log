package util

import (
	"fmt"
	"net/http"
	"strings"
)

// RenderTextPlain renders a slice of strings as text/plain; charset=utf-8
// If an error is passed in, it is rendered as the HTTP status code 500
func RenderTextPlain(w http.ResponseWriter, data []string, err error) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if err == nil {
		_, err = fmt.Fprintln(w, strings.Join(data, "\n"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = fmt.Fprintln(w, err.Error())
	}
}
