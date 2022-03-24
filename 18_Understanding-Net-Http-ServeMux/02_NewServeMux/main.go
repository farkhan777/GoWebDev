package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (m hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat\n")
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":7777", mux)
}
