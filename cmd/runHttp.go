package main

import (
	"fmt"
	"net/http"

	"github.com/runozo/golang-stdlib-http-app/router"
)

// runHttp runs an HTTP server on the specified listen address.
//
// listenAddr: The address to listen on.
// Returns an error if there was a problem starting the server.
func runHttp(listenAddr string) error {
	s := http.Server{
		Addr:    listenAddr,
		Handler: router.NewRouter(),
	}
	fmt.Printf("Starting HTTP listener at %s\n", listenAddr)
	return s.ListenAndServe()
}
