package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", home)
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

func home(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(res, "something.gohtml", "Farkhan")
	if err != nil {
		log.Fatalln(err)
	}
}
