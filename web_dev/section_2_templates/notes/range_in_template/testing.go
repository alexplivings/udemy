package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	stringBois := []string{"Hello", "you", "stupid", "person"}

	err := tpl.ExecuteTemplate(os.Stdout, "basic.gohtml", stringBois)
	if err != nil {
		log.Fatal(err)
	}
}
