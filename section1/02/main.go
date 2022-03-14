package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

func init() {
	tpl = template.Must(template.ParseFiles("dmg.gohtml"))
}

type hotels []hotel

func main() {

	xh := hotels{
		hotel{
			Name:    "Holiday Inn",
			Address: "123 makebelieve lane",
			City:    "San Francisco",
			Zip:     "12345",
			Region:  "Southern",
		},
		hotel{
			Name:    "Embassy Suites",
			Address: "99999 Dont know where",
			City:    "Napa",
			Zip:     "88888",
			Region:  "Northern",
		},
		hotel{
			Name:    "Hotel California",
			Address: "111 You can never leave",
			City:    "San Diego",
			Zip:     "99986",
			Region:  "Southern",
		},
	}

	err := tpl.Execute(os.Stdout, xh)

	if err != nil {
		log.Fatalln((err))
	}
}
