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
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		forum.LoginHandler(w, r)
	})
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		forum.LogoutHandler(w, r)
	})
	http.HandleFunc("/like_dislike", func(w http.ResponseWriter, r *http.Request) {
		forum.LikeHandler(w, r)
	})
	http.HandleFunc("/postcreation", func(w http.ResponseWriter, r *http.Request) {
		forum.CreationPost(w, r)
	})
	http.HandleFunc("/commentcreation", func(w http.ResponseWriter, r *http.Request) {
		forum.PublicationComment(w, r)
	})
	http.HandleFunc("/comment", func(w http.ResponseWriter, r *http.Request) {
		forum.MoreComment(w, r)
	})
	http.HandleFunc("/activity", func(w http.ResponseWriter, r *http.Request) {
		forum.Activity(w, r)
	})
	http.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
		forum.RemovePostHandler(w, r)
	})
	http.HandleFunc("/removeComm", func(w http.ResponseWriter, r *http.Request) {
		forum.RemoveCommHandler(w, r)
	})
	http.HandleFunc("/edit", func(w http.ResponseWriter, r *http.Request) {
		forum.EditPostHandler(w, r)
	})
	http.HandleFunc("/editComm", func(w http.ResponseWriter, r *http.Request) {
		forum.EditCommHandler(w, r)
	})
	http.HandleFunc("/removePic", func(w http.ResponseWriter, r *http.Request) {
		forum.RemovePicHandler(w, r)
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fs2 := http.FileServer(http.Dir("img/"))
	http.Handle("/img/", http.StripPrefix("/img/", fs2))

	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
