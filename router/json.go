/* router/json.go */
package router

import (
	"encoding/json"
	"net/http"
)

// JSON serializes the given data as JSON and writes it to the ResponseWriter.
//
// It takes three parameters:
//   - w: the http.ResponseWriter to write the JSON response to.
//   - data: the data to be serialized as JSON.
//   - code: the HTTP status code to be written in the response header.
//
// It does not return any value.
func JSON(w http.ResponseWriter, data interface{}, code int) {
	buf, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	status := http.StatusOK
	if code != 0 {
		status = code
	}
	w.WriteHeader(status)
	w.Write(buf)
}
