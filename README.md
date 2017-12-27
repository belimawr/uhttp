µhttp
=====

**µhttp** (micro-http) is a minimal http-framework or just a easier way to
handle Golang http.Handler errors in a easier way.

Motivation
----------
In all my Go project I always end up writing a thin wrapper over http.Handler
that can return a error and using it to avoid duplicating code to set HTTP
status code, error messages, default headers, etc.

So I decided to write something, a wee library to centralise all those ideas
and implementations

TODO
----

* Finish it :P
* Write unit-tests
* Good documentation
* Examples
* Contributors notice/guide/file

Licence
=======
MIT Licence
