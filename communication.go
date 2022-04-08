package forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

func CreationPost(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/postcreation.html"))

	switch r.Method {
	case "GET":
	case "POST":

		ck_uuid_user, err := r.Cookie("uuid_hash")
		if err != nil {
			fmt.Println(err)
		}
		uuid_user := ck_uuid_user.Value
		r.ParseForm()

		// Je récupère les valeurs de l'utilisateur
		title := r.FormValue("title")
		content := r.FormValue("content")

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
			category := "sport"
			// J'envoie les informations dans la base de donnée
			_, err = db.Exec("INSERT INTO posts(id, id_account, title, picture_text, category) VALUES(?,?,?,?,?)", id, uuid_user, title, content, category)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	data := ""
	tmpl.Execute(w, data)
}

func PublicationComment(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/commentcreation.html"))

	switch r.Method {
	case "GET":
	case "POST":
		// Je récupère la valeur de l'utilisateur
		content := r.FormValue("content")

		// Cela permet d'ouvrir et fermer la database
		db, err := sql.Open("sqlite3", "./forum.db")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		// Je récupère l'UUID de la personne pour prouver que ce commentaire est bien à lui
		//rows, err := db.Query("SELECT UUID FROM authentication WHERE UUID=(?);", uuid_user)

		// Je récupère l'ID du poste pour montrer que le commentaire est bien relier au poste
		//post, err := db.Query("SELECT INTO posts(id) WHERE id =(?)" /*variable id post*/)

		// J'envoie les informations dans la base de donnée
		_, err = db.Exec("INSERT INTO comments(comment) VALUES(?)", content)
		if err != nil {
			fmt.Println(err)
		}
	}
	data := ""
	tmpl.Execute(w, data)
}
