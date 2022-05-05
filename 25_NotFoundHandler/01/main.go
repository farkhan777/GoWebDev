package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":7777", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	fmt.Println(w, "go look at your terminal")
}
