package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/name/", name)

	http.ListenAndServe(":7777", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog")
}

func name(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Farkhan Hamzah Firdaus")
}
