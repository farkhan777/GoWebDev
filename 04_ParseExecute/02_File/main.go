package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("tpl.footerstuff.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.footerstuff.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// Do not use the above code in production
