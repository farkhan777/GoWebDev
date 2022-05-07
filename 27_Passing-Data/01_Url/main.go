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
	v := req.FormValue("q")
	fmt.Fprintln(w, "Do my search: "+v)
}

// visit this page:
// http://localhost:7777/?q=dog
