package main

import (
	"fmt"
	"forum"
	"net/http"
)

type ToDoPage struct {
	PageTitle string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		forum.MainHandler(w, r)
	})
	http.HandleFunc("/authentication", func(w http.ResponseWriter, r *http.Request) {
		forum.AuthenticationHandler(w, r)
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
