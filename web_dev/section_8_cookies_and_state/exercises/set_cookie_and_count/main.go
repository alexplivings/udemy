package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "1",
		Path:  "/",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	count, err := strconv.ParseInt(c.Value, 0, 64)
	if err != nil {
		fmt.Printf("Error in parsing counter: %v", err)
	}
	count++

	fmt.Printf("Counter value: %v\n", count)

	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: fmt.Sprintf("%v", count),
		Path:  "/",
	})

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}
