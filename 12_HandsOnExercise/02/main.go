package main

import (
	"log"
	"os"
	"text/template"
)

type region struct {
	Region string
	Hotels []hotel
}

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r1 := []region{
		{
			Region: "East Java",
			Hotels: []hotel{
				{"Ijen View", "Jln.Bondowoso", "Bondowoso", "6385"},
				{"Pal Hotel", "Jln.Raya Bondowoso", "Bondowoso", "6385"},
			},
		},
		{
			Region: "West Java",
			Hotels: []hotel{
				{"Crowne Plaza Bandung", "Jl. Lembong No.19, Braga, Kec. Sumur Bandung, Kota Bandung, Jawa Barat", "Bandung", "40111"},
				{"Zest Hotel Sukajadi", "Jl. Sukajadi No.16, Pasteur, Kec. Sukajadi, Kota Bandung, Jawa Barat", "Bandung", "40162"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, r1)
	if err != nil {
		log.Fatalln(err)
	}
}
