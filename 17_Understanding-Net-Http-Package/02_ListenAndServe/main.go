package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
	fmt.Fprintln(w, "Farkhan Hamzah Firdaus")
}

func main() {
	var d hotdog
	http.ListenAndServe(":7777", d)
}
