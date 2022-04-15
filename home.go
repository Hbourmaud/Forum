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
	Picture_text       string
	Title              string
	Category           string
	Like               int
	Dislike            int
	One_comment        string
	One_comment_author string
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
	var Picture_textTab []string
	var TitleTab []string
	var CategoryTab []string
	var LikeTab []int
	var DislikeTab []int
	var One_commentTab []string
	var One_comment_authorTab []string

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
		One_comment_authorTab = append(One_comment_authorTab, "")
		One_commentTab = append(One_commentTab, "")

		dataComment, err := db.Query("SELECT id_account, comment FROM comments WHERE id_post=(?) ORDER BY id DESC LIMIT 1", post.Id_post)
		if err != nil {
			fmt.Println(err)
		}
		defer dataComment.Close()

		for dataComment.Next() {
			var Id_comment_author string
			var Comment string
			err = dataComment.Scan(&Id_comment_author, &Comment)

			if err != nil {
				fmt.Println(err)
			}
			One_comment_authorTab[i] = Id_comment_author
			One_commentTab[i] = Comment
		}

		dataLike, err2 := db.Query("SELECT count() FROM likes WHERE id_post=(?)", post.Id_post)
		if err2 != nil {
			fmt.Println(err)
		}
		defer dataPost.Close()
		for dataLike.Next() {
			var like_post int
			err3 := dataLike.Scan(&like_post)
			if err3 != nil {
				fmt.Println(err3)
			}
			LikeTab = append(LikeTab, like_post)
		}
		dataDislike, err4 := db.Query("SELECT count() FROM dislikes WHERE id_post=(?)", post.Id_post)
		if err4 != nil {
			fmt.Println(err)
		}
		defer dataDislike.Close()
		for dataDislike.Next() {
			var dislike_post int
			err5 := dataDislike.Scan(&dislike_post)
			if err5 != nil {
				fmt.Println(err5)
			}
			DislikeTab = append(DislikeTab, dislike_post)
		}
		post.One_comment = One_commentTab[i]
		post.One_comment_author = One_comment_authorTab[i]
		post.Like = LikeTab[i]
		post.Dislike = DislikeTab[i]
		DataTab = append(DataTab, post)
	}

	switch r.Method {
	case "POST":
		category := r.FormValue("category")
		DataTab = filters(category)
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
