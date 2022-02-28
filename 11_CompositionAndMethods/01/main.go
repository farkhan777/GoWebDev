package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("hw.gohtml"))
}

func main() {
	p1 := person{
		Name: "Farkhan Hamzah Firdaus",
		Age:  21,
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
