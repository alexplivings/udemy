package main

import (
	"fmt"
	"net/http"
)

func rootHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "root")
}

func dogHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "dog")
}

func meHandle(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Alex Livingston")
}

func main() {

	http.HandleFunc("/", rootHandle)
	http.HandleFunc("/dog/", dogHandle)
	http.HandleFunc("/me/", meHandle)

	http.ListenAndServe(":8080", nil)

}
