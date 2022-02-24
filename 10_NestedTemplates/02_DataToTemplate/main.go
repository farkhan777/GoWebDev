package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*footerstuff.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.footerstuff.gohtml", 100)
	if err != nil {
		log.Fatalln(err)
	}
}
