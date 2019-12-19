package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("root.gohtml", "dog.gohtml", "me.gohtml"))

func rootHandle(res http.ResponseWriter, req *http.Request) {
	var err = templates.ExecuteTemplate(res, "root.gohtml", nil)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func dogHandle(res http.ResponseWriter, req *http.Request) {
	var err = templates.ExecuteTemplate(res, "dog.gohtml", nil)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func meHandle(res http.ResponseWriter, req *http.Request) {
	data := "Alex Livingston"
	var err = templates.ExecuteTemplate(res, "me.gohtml", data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.HandleFunc("/", rootHandle)
	http.HandleFunc("/dog/", dogHandle)
	http.HandleFunc("/me/", meHandle)

	http.ListenAndServe(":8080", nil)

}
