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
	Id_post            int
	Id_account         string
	Texts              string
	Title              string
	Category           string
	Picture            string
	Like               int
	Dislike            int
	One_comment        string
	One_comment_author string
	One_comment_id     string
	Comments           []CommentStruct
}

type CommentStruct struct {
	Id_post         int
	Id_comment      int
	Id_account      string
	Comment         string
	Like_comment    int
	Dislike_comment int
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))

	// variable post
	var Id_postTab []int       /* réutiliser pour le commentaire */
	var Id_accountTab []string /* réutiliser pour le commentaire */
	var TextsTab []string
	var TitleTab []string
	var CategoryTab []string
	var PictureTab []string

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
		var texts string
		var Title string
		var category string
		var picture sql.NullString

		err = dataPost.Scan(&id_post, &id_account, &Title, &texts, &category, &picture)

		if err != nil {
			fmt.Println(err)
		}

		Id_postTab = append(Id_postTab, id_post)
		Id_accountTab = append(Id_accountTab, id_account)
		TextsTab = append(TextsTab, texts)
		TitleTab = append(TitleTab, Title)
		CategoryTab = append(CategoryTab, category)
		PictureTab = append(PictureTab, picture.String)

	}
	var post Info
	for i := 0; i < len(Id_postTab); i++ {
		post.Id_post = Id_postTab[i]
		post.Id_account = idAccount_to_username(Id_accountTab[i])
		post.Title = TitleTab[i]
		post.Category = CategoryTab[i]
		post.Texts = TextsTab[i]
		post.Picture = PictureTab[i]

		DataTab = append(DataTab, post)
	}
	DataTab = addDetailsPost(DataTab)

	switch r.Method {
	case "POST":
		category := r.FormValue("category")
		user_uuid, err := r.Cookie("uuid_hash")
		uuid_user := ""
		if err != nil {
			if category == "" {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			fmt.Println(err)
		} else {
			uuid_user = user_uuid.Value
		}
		created_posted := r.FormValue("created_posts")
		liked_posts := r.FormValue("liked_posts")
		DataTab = filters(category, uuid_user, created_posted, liked_posts)
		DataTab = addDetailsPost(DataTab)
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
