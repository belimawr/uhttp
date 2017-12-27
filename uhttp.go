// Package uhttp is a minimal http-framework or just a easier way to
// handle Golang http.Handler errors in a easier way
package uhttp

import "net/http"

// Handler interface to describe an http.Handler that returns an error
type Handler interface {
	// ServeHTTP http.Handler that returns an error
	ServeHTTP(w http.ResponseWriter, r *http.Request) error
}

// HandlerFunc wrapper to allow functions to be used as Handler
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

// HTTPError describes an error that holds: HTTP status code,
// user-friendly message, headers and also the error cause
type HTTPError interface {
	error

	// HTTP status code to be set
	StatusCode() int

	// Body that can be sent to the final user of the API
	UserFriendlyError() []byte

	// Headers to be set in an error occurs
	Headers() map[string]string

	// Cause of the error
	Cause() error
}

// ServeHTTP allows HandlerFunc to be used as Handler
func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	return h(w, r)
}
