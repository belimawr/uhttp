µhttp
=====

**µhttp** (micro-http) is a minimal http-framework or just a easier way to
handle Golang http.Handler errors in a easier way.

Motivation
----------
The standard net/http Handler does not provide a way to handle errors, the code
of a handler is full of:

```go
if errorCondition {
	w.Header().Set("X-Error-Header", "sorry, there was an error")
   	w.WriteHeader(http.StatusBadRequest)
	w.Write("some error message")
	return
}
```

To avoid this in all my Go project I always end up writing a thin wrapper
around http.Handler that can return a error and use it to avoid duplicating
the code to set HTTP status code, write headers, error messages, etc.

So I decided to write a wee library to centralise all those ideas
and implementations

Usage
-----
```sh
go get github.com/belimawr/uhttp
```

Or if you are using [dep](https://github.com/golang/dep)
```sh
dep ensure -add github.com/belimawr/uhttp
```

Then you just need to write your handlers:
```go
func myHandler(w http.ResponseWriter, r *http.Request) error {
	if err := somethingThatCanGoWrong(r); err != nil{
		return err
	}
}
```

And wrap them
```go
http.Handle("/foo", uhttp.WrapFunc(myHandler))
```

The default status code to errors is ``http.StatusInternalServerError``,
however if the returned error implements the interface ```HTTPError```
(see [uhttp.go](https://github.com/belimawr/uhttp/blob/master/uhttp.go#L18-L32))
then the status code and the headers of the error will be used.
See [examples/main.go](examples/main.go) for a full example.

TODO
----

* Write unit-tests
* Improve documentation
* Improve Examples
* Contributors notice/guide/file

Licence
=======
MIT Licence
