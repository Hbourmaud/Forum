package forum

import (
	"database/sql"
	"fmt"
)

func addDetailsPost(DataTab []Info) []Info {
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	var LikeTab []int
	var DislikeTab []int
	var One_commentTab []string
	var One_comment_authorTab []string
	var One_comment_idTab []string
	var DataTabFinal []Info
	var post Info
	for i := 0; i < len(DataTab); i++ {
		post.Id_post = DataTab[i].Id_post
		post.Id_account = DataTab[i].Id_account
		post.Title = DataTab[i].Title
		post.Category = DataTab[i].Category
		post.Texts = DataTab[i].Texts
		post.Picture = DataTab[i].Picture
		One_comment_authorTab = append(One_comment_authorTab, "")
		One_commentTab = append(One_commentTab, "")
		One_comment_idTab = append(One_comment_idTab, "no")

		dataComment, err := db.Query("SELECT id,id_account, comment FROM comments WHERE id_post=(?) ORDER BY id DESC LIMIT 1", post.Id_post)
		if err != nil {
			fmt.Println(err)
		}
		defer dataComment.Close()

		for dataComment.Next() {
			var Id_comment_author string
			var Comment string
			var Comment_id string
			err = dataComment.Scan(&Comment_id, &Id_comment_author, &Comment)

			if err != nil {
				fmt.Println(err)
			}
			One_comment_authorTab[i] = Id_comment_author
			One_commentTab[i] = Comment
			One_comment_idTab[i] = Comment_id

		}

		dataLike, err2 := db.Query("SELECT count() FROM likes WHERE id_post=(?)", post.Id_post)
		if err2 != nil {
			fmt.Println(err)
		}
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
		post.One_comment_author = idAccount_to_username(One_comment_authorTab[i])
		post.One_comment_id = One_comment_idTab[i]
		post.Like = LikeTab[i]
		post.Dislike = DislikeTab[i]
		DataTabFinal = append(DataTabFinal, post)
	}
	return DataTabFinal
}

func idAccount_to_username(id_account string) string {
	if id_account == "" {
		return ""
	}
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	dataUserName, err := db.Query("SELECT username FROM authentication WHERE UUID=(?)", id_account)
	if err != nil {
		fmt.Println("error", err)
	}
	defer dataUserName.Close()
	var username string
	for dataUserName.Next() {
		err = dataUserName.Scan(&username)
		if err != nil {
			fmt.Println(err)
		}
	}
	return username
}
