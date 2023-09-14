package router

import (
	"fmt"
	"net/http"
)

/* router/errors.go */
type httpErr struct {
	err error
	msg string
}

// Error returns the string representation of the httpErr error.
//
// No parameters.
// Returns a string.
func (he httpErr) Error() string {
	return he.err.Error()
}

// ErrJSON handles JSON errors and writes them to the response writer.
//
// It takes three parameters:
// - w: the http.ResponseWriter used to write the error response.
// - err: the error that occurred.
// - code: the HTTP status code to be used in the response.
//
// This function does not return anything.
func ErrJSON(w http.ResponseWriter, err error, code int) {
	v, ok := err.(httpErr)
	if !ok {
		fmt.Printf("An error occurred: %v", err)
		http.Error(w, err.Error(), code)
		return
	}
	fmt.Printf("An error occurred: %v", v.err)
	JSON(w, map[string]string{
		"error": v.msg,
	}, code)
}
