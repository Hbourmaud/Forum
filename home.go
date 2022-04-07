package forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

type Data struct {
	DataTab   []Info
	Username  string
	UUID_hash string
}

type Info struct {
	Id_post      int
	Id_account   string
	Picture_text string
	Title        string
	Category     string
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))

	var Id_postTab []int
	var Id_accountTab []string
	var Picture_textTab []string
	var TitleTab []string
	var CategoryTab []string

	//Cela permet d'ouvrir et de fermer la database
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// Je récupère les informations des posts
	dataPost, err := db.Query("SELECT * FROM posts")
	if err != nil {
		fmt.Println(err)
	}
	defer dataPost.Close()
	var DataTab []Info
	username := "not login"
	UUID_string := ""
	for dataPost.Next() {
		var id_post int
		var id_account string
		var picture_text string
		var Title string
		var category string

		err = dataPost.Scan(&id_post, &id_account, &picture_text, &Title, &category)

		if err != nil {
			fmt.Println(err)
		}

		Id_postTab = append(Id_postTab, id_post)
		Id_accountTab = append(Id_accountTab, id_account)
		Picture_textTab = append(Picture_textTab, picture_text)
		TitleTab = append(TitleTab, Title)
		CategoryTab = append(CategoryTab, category)
	}
	var post Info
	for i := 0; i < len(Id_postTab); i++ {
		post.Id_post = Id_postTab[i]
		post.Id_account = Id_accountTab[i]
		post.Title = TitleTab[i]
		post.Category = CategoryTab[i]
		post.Picture_text = Picture_textTab[i]

		DataTab = append(DataTab, post)
	}
	ck_user, err := r.Cookie("username")
	if err != nil {
		fmt.Println(err)
	}
	ck_uuid, err := r.Cookie("uuid_hash")
	if err != nil {
		fmt.Println(err)
	}
	if ck_user != nil && ck_uuid != nil {
		username = ck_user.Value
		UUID_string = ck_uuid.Value
	}
	data := &Data{
		DataTab:   DataTab,
		Username:  username,
		UUID_hash: UUID_string,
	}
	err2 := tmpl.Execute(w, data)
	if err2 != nil {
		fmt.Println(err2)
	}
}
