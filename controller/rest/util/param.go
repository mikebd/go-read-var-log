package util

import (
	"fmt"
	"net/http"
	"strconv"
)

// PositiveIntParamStrict returns the value of the parameter as an int (>= 0), or the default value if the parameter is not present
func PositiveIntParamStrict(w http.ResponseWriter, r *http.Request, defaultValue int, param string) (int, error) {
	paramValue := r.URL.Query().Get(param)
	if paramValue == "" {
		return defaultValue, nil
	}

	intParam, err := strconv.Atoi(paramValue)
	if intParam < 0 || err != nil {
		RenderTextPlain(w, nil, fmt.Errorf("invalid value for parameter '%s': '%s'", param, paramValue))
		return 0, err
	}

	return intParam, nil
}
