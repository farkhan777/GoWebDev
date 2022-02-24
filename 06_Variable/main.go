package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.footerstuff.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.footerstuff.gohtml", `Release self-focus; embrace other-focus.`)
	if err != nil {
		log.Fatalln(err)
	}
}
