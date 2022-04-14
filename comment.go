package forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func MoreComment(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/comment.html"))

	switch r.Method {
	case "GET":
	case "POST":

		uuid_post := r.FormValue("uuid_post")

		// variable post
		var Id_post int
		var Id_account string
		var Picture_text string
		var Title string
		var Category string

		//variable commentaire
		var Id_commentTab []int
		var Id_accountTab []string
		var CommentTab []string

		//Cela permet d'ouvrir et de fermer la database
		db, err := sql.Open("sqlite3", "./forum.db")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		// Je récupère les informations des posts
		dataPost, err := db.Query("SELECT * FROM posts WHERE id =(?)", uuid_post)
		if err != nil {
			fmt.Println(err)
		}
		defer dataPost.Close()

		var DComment []CommentStruct
		for dataPost.Next() {
			var id_post int
			var id_account string
			var picture_text string
			var title string
			var category string

			err = dataPost.Scan(&id_post, &id_account, &picture_text, &Title, &category)

			if err != nil {
				fmt.Println(err)
			}

			Id_post = id_post
			Id_account = id_account
			Picture_text = picture_text
			Title = title
			Category = category
		}

		var post Info

		post.Id_post = Id_post
		post.Id_account = Id_account
		post.Title = Title
		post.Category = Category
		post.Picture_text = Picture_text

		dataComment, err := db.Query("SELECT id, id_account, comment FROM comments WHERE id_post=(?)", uuid_post)
		if err != nil {
			fmt.Println(err)
		}
		defer dataComment.Close()

		for dataComment.Next() {
			var Id_comment int
			var Id_account string
			var Comment string

			err = dataComment.Scan(&Id_comment, &Id_account, &Comment)

			if err != nil {
				fmt.Println(err)
			}

			Id_commentTab = append(Id_commentTab, Id_comment)
			Id_accountTab = append(Id_accountTab, Id_account)
			CommentTab = append(CommentTab, Comment)

			var comment CommentStruct
			for i := 0; i < len(Id_commentTab); i++ {
				comment.Id_comment = Id_commentTab[i]
				comment.Id_account = Id_accountTab[i]
				comment.Comment = CommentTab[i]
			}
			DComment = append(DComment, comment)
		}
		post.Comments = DComment

		int_uuid_post, err1 := strconv.Atoi(uuid_post)
		if err1 != nil {
			fmt.Println(err)
		}

		data := &Info{
			Id_post:      int_uuid_post,
			Id_account:   Id_account,
			Picture_text: Picture_text,
			Title:        Title,
			Category:     Category,
			Comments:     DComment,
		}
		err2 := tmpl.Execute(w, data)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
}
