package main

import (
	"errors"
	"net/http"

	"github.com/belimawr/uhttp"
)

func NewBadRequestError(e error) uhttp.HTTPError {
	return uhttp.NewHTTPError(e,
		http.StatusBadRequest,
		[]byte(`Bad request, you do not know what you are doing!`),
		map[string]string{})
}

func myHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return NewBadRequestError(errors.New("Invalid Method"))
	}

	_, err := w.Write([]byte("Hello World!"))
	return err
}

func main() {
	http.Handle("/hello", uhttp.WrapFunc(myHandler))

	panic(http.ListenAndServe(":3000", nil))
}
