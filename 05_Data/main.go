package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.footerstuff.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.footerstuff.gohtml", 100)
	if err != nil {
		log.Fatalln(err)
	}
}
