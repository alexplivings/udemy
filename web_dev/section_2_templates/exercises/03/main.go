package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []hotel{
		hotel{"Test", "123 Test Rd", "Testburg", "12345", "Test Valley"},
		hotel{"Boii", "123 Boi Rd", "Boiburg", "12345", "Boi Valley"},
		hotel{"Gurll", "123 Gurl Rd", "Gurlburg", "12345", "Gurl Valley"},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", hotels)
	if err != nil {
		log.Fatal(err)
	}
}

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}
