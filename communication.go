package forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

type PostIdStruct struct {
	Id_post string
}

func CreationPost(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/postcreation.html"))

	switch r.Method {
	case "GET":

	case "POST":
		ck_uuid_user, err := r.Cookie("uuid_hash")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			fmt.Println(err)
			return
		}
		uuid_user := ck_uuid_user.Value
		r.ParseForm()

		title, content, category, picture := uploadImage(w, r)

		if title == "" {

		} else {
			// Cela permet d'ouvrir et fermer la database
			db, err := sql.Open("sqlite3", "./forum.db")
			if err != nil {
				fmt.Println(err)
			}
			defer db.Close()

			// Je récupère l'UUID de la personne pour prouver que ce poste est bien à lui
			rows, err := db.Query("SELECT UUID FROM authentication WHERE UUID=(?);", uuid_user)
			if err != nil {
				fmt.Println(err)
			}

			for rows.Next() {
				var id_account int

				err = rows.Scan(&id_account)

				if err != nil {
					fmt.Println(err)
				}
			}
			var id sql.NullInt64
			// J'envoie les informations dans la base de donnée
			_, err = db.Exec("INSERT INTO posts(id, id_account, title, texts, category, picture) VALUES(?,?,?,?,?,?)", id, uuid_user, title, content, category, picture)
			if err != nil {
				fmt.Println(err)
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
	data := ""
	tmpl.Execute(w, data)
}

func PublicationComment(w http.ResponseWriter, r *http.Request) {
	var uuid_post string
	tmpl := template.Must(template.ParseFiles("static/commentcreation.html"))

	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case "POST":

		ck_uuid_user, err := r.Cookie("uuid_hash")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			fmt.Println(err)
			return
		}
		uuid_user := ck_uuid_user.Value
		r.ParseForm()

		// Je récupère la valeur de l'utilisateur
		content := r.FormValue("content")
		uuid_post = r.FormValue("uuid_post")

		if content == "" {

		} else {
			// Cela permet d'ouvrir et fermer la database
			db, err := sql.Open("sqlite3", "./forum.db")
			if err != nil {
				fmt.Println(err)
			}
			defer db.Close()

			// Je récupère l'UUID de la personne pour prouver que ce commentaire est bien à lui
			rows1, err := db.Query("SELECT UUID FROM authentication WHERE UUID=(?);", uuid_user)
			if err != nil {
				fmt.Println(err)
			}

			// Je récupère l'ID du poste pour montrer que le commentaire est bien relier au poste
			rows2, err := db.Query("SELECT id FROM posts WHERE id =(?)", uuid_post)
			if err != nil {
				fmt.Println(err)
			}

			for rows1.Next() {
				var id_account int

				err = rows1.Scan(&id_account)

				if err != nil {
					fmt.Println(err)
				}
			}

			for rows2.Next() {
				var id_post int

				err = rows2.Scan(&id_post)

				if err != nil {
					fmt.Println(err)
				}
			}
			var id sql.NullInt64
			// J'envoie les informations dans la base de donnée
			_, err = db.Exec("INSERT INTO comments(id, id_post, id_account, comment) VALUES(?,?,?,?)", id, uuid_post, uuid_user, content)
			if err != nil {
				fmt.Println(err)
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
	data := &PostIdStruct{
		Id_post: uuid_post,
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}
