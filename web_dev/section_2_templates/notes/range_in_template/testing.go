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

	err := tpl.ExecuteTemplate(os.Stdout, "with_index.gohtml", stringBois)
	if err != nil {
		log.Fatal(err)
	}
}
