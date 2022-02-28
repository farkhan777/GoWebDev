package main

import (
	"log"
	"os"
	"text/template"
)

type restaurant struct {
	Name  string
	Menus []menu
}

type menu struct {
	Meal  string
	Items []item
}

type item struct {
	Name        string
	Description string
	Price       float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m1 := []restaurant{
		{
			Name: "Federicos",
			Menus: []menu{
				{
					Meal: "Breakfast",
					Items: []item{
						{"Outmeal", "yum yum", 6.45},
						{"Cheerios", "American eating food traditional now", 3.95},
						{"Juice Orange", "Delicious drinking in throat squeezed fresh", 2.95},
					},
				},
				{
					Meal: "Lunch",
					Items: []item{
						item{
							Name:        "Hamburger",
							Description: "Delicous good eating for you",
							Price:       6.95,
						},
						item{
							Name:        "Cheese Melted Sandwich",
							Description: "Make cheese bread melt grease hot",
							Price:       3.95,
						},
						item{
							Name:        "French Fries",
							Description: "French eat potatoe fingers",
							Price:       2.95,
						},
					},
				},
				{
					Meal: "Dinner",
					Items: []item{
						item{
							Name:        "Pasta Bolognese",
							Description: "From Italy delicious eating",
							Price:       7.95,
						},
						item{
							Name:        "Steak",
							Description: "Dead cow grilled bloody",
							Price:       13.95,
						},
						item{
							Name:        "Bistro Potatoe",
							Description: "Bistro bar wood American bacon",
							Price:       6.95,
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m1)
	if err != nil {
		log.Fatalln(err)
	}
}
