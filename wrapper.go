package uhttp

import (
	"log"
	"net/http"
)

// Wrap wraps an uhttp.Handler making it an http.Handler.
//
// Errors are handled properly by setting:
//  * status conde
//  * headers
//  * body
func Wrap(h Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h.ServeHTTP(w, r)

		if err == nil {
			return
		}

		switch e := err.(type) {
		case HTTPError:
			for k, v := range e.Headers() {
				w.Header().Add(k, v)
			}

			w.WriteHeader(e.StatusCode())
			w.Write(e.UserFriendlyError())

			// Golang should have a Logger interface...
			// TODO: find a more flexible way to log the error
			log.Print(e.Error())
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Seever Error"))

			log.Println(err)
		}
	})
}

// WrapFunc wraps an uhttp.HandlerFunc making it an http.Handler
// see Wrapper for more details
func WrapFunc(h HandlerFunc) http.Handler {
	return Wrap(HandlerFunc(h))
}
