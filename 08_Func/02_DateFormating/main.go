package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.footerstuff.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("02_SQL-01_Rot13-2006")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.footerstuff.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
