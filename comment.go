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
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case "POST":

		uuid_post := r.FormValue("uuid_post")

		// variable post
		var id_post int
		var id_account string
		var texts string
		var title string
		var category string
		var picture string
		//variable commentaire
		var Id_commentTab []int
		var Id_accountTab []string
		var CommentTab []string
		var Like_commentTab []int
		var Dislike_commentTab []int

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

			err = dataPost.Scan(&id_post, &id_account, &title, &texts, &category, &picture)

			if err != nil {
				fmt.Println(err)
			}

		}

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

		for i := 0; i < len(Id_commentTab); i++ {
			dataLike, err2 := db.Query("SELECT count() FROM likes WHERE id_comment=(?)", Id_commentTab[i])
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
				Like_commentTab = append(Like_commentTab, like_post)
			}
			dataDislike, err4 := db.Query("SELECT count() FROM dislikes WHERE id_comment=(?)", Id_commentTab[i])
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
				Dislike_commentTab = append(Dislike_commentTab, dislike_post)
			}
			DComment[i].Like_comment = Like_commentTab[i]
			DComment[i].Dislike_comment = Dislike_commentTab[i]
		}

		int_uuid_post, err1 := strconv.Atoi(uuid_post)
		if err1 != nil {
			fmt.Println(err)
		}

		data := &Info{
			Id_post:    int_uuid_post,
			Id_account: id_account,
			Texts:      texts,
			Title:      title,
			Category:   category,
			Comments:   DComment,
		}
		err2 := tmpl.Execute(w, data)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
}
