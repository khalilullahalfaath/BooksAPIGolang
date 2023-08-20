package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody is a utility function to parse the request body
// and return the body as a map
func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
