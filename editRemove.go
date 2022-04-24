package forum

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type EditInfo struct {
	Id_post  string
	Texts    string
	Title    string
	Category string
	Picture  string
	Sentence string
	IsGood   string
}

type EditCommStruct struct {
	Id_comm string
	Comment string
}

func RemovePostHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("uuid_hash")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		fmt.Println(err)
		return
	}
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case "POST":
		body, _ := ioutil.ReadAll(r.Body)
		keyVal := make(map[string]string)
		json.Unmarshal(body, &keyVal)
		id_post := keyVal["id_post"]
		db, err := sql.Open("sqlite3", "./forum.db")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
		rows, err3 := db.Query("SELECT picture FROM posts WHERE id=(?);", id_post)
		if err != nil {
			fmt.Println(err3)
			return
		}
		for rows.Next() {
			var path_pictureDel string
			err = rows.Scan(&path_pictureDel)
			if err != nil {
				fmt.Println(err)
			}
			if path_pictureDel != "" {
				err = os.Remove(path_pictureDel)
				if err != nil {
					fmt.Println(err)
				}
			}

		}
		rows.Close()
		_, err2 := db.Exec("DELETE FROM posts WHERE id = (?);", id_post)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		_, err2 = db.Exec("DELETE FROM likes WHERE id_post = (?);", id_post)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		_, err2 = db.Exec("DELETE FROM dislikes WHERE id_post = (?);", id_post)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		rows2, err2 := db.Query("SELECT id FROM comments WHERE id_post=(?);", id_post)
		if err != nil {
			fmt.Println(err2)
			return
		}
		var commentToDestroy []int
		for rows2.Next() {
			var id_comment int
			err = rows2.Scan(&id_comment)
			if err != nil {
				fmt.Println(err)
			}
			commentToDestroy = append(commentToDestroy, id_comment)
		}
		defer rows2.Close()
		for _, id_comment_del := range commentToDestroy {
			_, err2 = db.Exec("DELETE FROM likes WHERE id_comment = (?);", id_comment_del)
			if err2 != nil {
				fmt.Println(err2)
			}
			_, err2 = db.Exec("DELETE FROM dislikes WHERE id_comment = (?);", id_comment_del)
			if err2 != nil {
				fmt.Println(err2)
			}
		}
		_, err2 = db.Exec("DELETE FROM comments WHERE id_post = (?);", id_post)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
	}
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/edit.html"))
	var id_post string
	var title string
	var texts string
	var content string
	var category string
	var picture string
	var sentence string
	var isGood string
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case "POST":
		id_post = r.FormValue("edit_id_post")
		if id_post == "" {
			id_post = r.FormValue("post_to_change")
		}
		statusEdit := r.FormValue("status_edit")
		db, err := sql.Open("sqlite3", "./forum.db")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
		rows, err := db.Query("SELECT title,texts,category,picture FROM posts WHERE id =(?)", id_post)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&title, &texts, &category, &picture)
			if err != nil {
				fmt.Println(err)
			}
		}
		if statusEdit != "toEdit" {
			temp_title, temp_content, temp_category, temp_picture := uploadImage(w, r)
			if temp_picture == "" && temp_title == "" {

			} else if temp_picture == "tooBig" {
				sentence = " File is too big, Maximum 20mb file accepted"
				isGood = "toEdit"
			} else if temp_picture == "wrongType" {
				sentence = " Wrong type file, please upload JPG, PNG or GIF"
				isGood = "toEdit"
			} else {
				title, content, category, picture = temp_title, temp_content, temp_category, temp_picture
				err := r.ParseMultipartForm(Max_upload_size)
				if err != nil {
					fmt.Println(err)
				}
				id_post = r.FormValue("post_to_change")
				_, err = db.Exec("UPDATE posts SET title = (?), texts = (?), category = (?), picture = (?) WHERE id=(?);", title, content, category, picture, id_post)
				if err != nil {
					fmt.Println(err)
				}
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		}
	}
	data := &EditInfo{
		Id_post:  id_post,
		Texts:    texts,
		Title:    title,
		Category: category,
		Picture:  picture,
		Sentence: sentence,
		IsGood:   isGood,
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func RemovePicHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case "POST":
		body, _ := ioutil.ReadAll(r.Body)
		keyVal := make(map[string]string)
		json.Unmarshal(body, &keyVal)
		picture := keyVal["picture"]
		db, err := sql.Open("sqlite3", "./forum.db")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
		_, err1 := db.Exec("UPDATE posts set picture = '' WHERE picture=(?);", picture)
		if err1 != nil {
			fmt.Println(err)
		}
		err = os.Remove(picture)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func RemoveCommHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("uuid_hash")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		fmt.Println(err)
		return
	}
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case "POST":
		body, _ := ioutil.ReadAll(r.Body)
		keyVal := make(map[string]string)
		json.Unmarshal(body, &keyVal)
		id_comm := keyVal["id_comm"]
		db, err := sql.Open("sqlite3", "./forum.db")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
		_, err2 := db.Exec("DELETE FROM likes WHERE id_comment = (?);", id_comm)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		_, err2 = db.Exec("DELETE FROM dislikes WHERE id_comment = (?);", id_comm)
		if err2 != nil {
			fmt.Println(err2)
			return
		}
		_, err2 = db.Exec("DELETE FROM comments WHERE id=(?);", id_comm)
		if err != nil {
			fmt.Println(err2)
			return
		}
	}
}

func EditCommHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/editComm.html"))
	var id_comm string
	var comment string
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case "POST":
		id_comm = r.FormValue("edit_id_comm")
		statusEdit := r.FormValue("status_edit")
		db, err := sql.Open("sqlite3", "./forum.db")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
		rows, err := db.Query("SELECT comment FROM comments WHERE id =(?)", id_comm)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&comment)
			if err != nil {
				fmt.Println(err)
			}
		}
		if statusEdit != "toEdit" {
			comment = r.FormValue("comment")
			if err != nil {
				fmt.Println(err)
			}
			_, err = db.Exec("UPDATE comments SET comment = (?) WHERE id=(?);", comment, id_comm)
			if err != nil {
				fmt.Println(err)
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
	data := &EditCommStruct{
		Id_comm: id_comm,
		Comment: comment,
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}
