package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":7777", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request method at foo: ", r.Method, "\n\n")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at bar:", r.Method)
	// process from data
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at barred:", r.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
