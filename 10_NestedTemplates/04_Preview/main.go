package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*footerstuff.gohtml"))
}

func main() {
	http.HandleFunc("/", index)

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(res, "indexgohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
