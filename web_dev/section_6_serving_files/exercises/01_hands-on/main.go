package main

import (
	"io"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseFiles("dog.gohtml"))

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	var err = templates.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpg")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", dogPic)

	http.ListenAndServe(":8080", nil)
}
