package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody parses the request body into the provided struct
func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal((body), x); err != nil {
			return
		}
	}
}
