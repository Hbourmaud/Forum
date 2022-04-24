package forum

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func LikeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case "POST":
		body, _ := ioutil.ReadAll(r.Body)
		keyVal := make(map[string]string)
		json.Unmarshal(body, &keyVal)
		like_or_dislike := keyVal["like_or_dislike"]
		id_post := keyVal["id_post"]
		id_comment := keyVal["id_comment"]
		isPost := false
		isComment := false
		db, err := sql.Open("sqlite3", "./forum.db")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		ck_uuid_user, err2 := r.Cookie("uuid_hash")
		if err2 != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			fmt.Println(err2)
			return
		}
		uuid_user := ck_uuid_user.Value
		r.ParseForm()
		if id_post != "" {
			isPost = true
		} else {
			isComment = true
		}
		var temp_count_like int
		var temp_count_dislike int
		if isPost {
			rows, err3 := db.Query("SELECT count(*) FROM likes WHERE (id_account, id_post) =(?,?);", uuid_user, id_post)
			if err3 != nil {
				fmt.Println(err3)
			}
			defer rows.Close()
			for rows.Next() {
				err4 := rows.Scan(&temp_count_like)
				if err4 != nil {
					fmt.Println(err4)
				}
			}
			rows, err5 := db.Query("SELECT count(*) FROM dislikes WHERE (id_account, id_post) =(?,?);", uuid_user, id_post)
			if err5 != nil {
				fmt.Println(err5)
			}
			defer rows.Close()
			for rows.Next() {
				err6 := rows.Scan(&temp_count_dislike)
				if err6 != nil {
					fmt.Println(err6)
				}
			}
		} else if isComment {
			rows, err3 := db.Query("SELECT count(*) FROM likes WHERE (id_account, id_comment) =(?,?);", uuid_user, id_comment)
			if err3 != nil {
				fmt.Println(err3)
			}
			defer rows.Close()
			for rows.Next() {
				err4 := rows.Scan(&temp_count_like)
				if err4 != nil {
					fmt.Println(err4)
				}
			}
			rows, err5 := db.Query("SELECT count(*) FROM dislikes WHERE (id_account, id_comment) =(?,?);", uuid_user, id_comment)
			if err5 != nil {
				fmt.Println(err5)
			}
			defer rows.Close()
			for rows.Next() {
				err6 := rows.Scan(&temp_count_dislike)
				if err6 != nil {
					fmt.Println(err6)
				}
			}
		}
		if like_or_dislike == "like" {
			if isPost {
				if temp_count_like > 0 {
					_, err3 := db.Exec("DELETE FROM likes WHERE (id_account, id_post) = (?,?);", uuid_user, id_post)
					if err3 != nil {
						fmt.Println(err)
					}
					fmt.Println(" removed post liked by reclick")
				} else {
					_, err2 := db.Exec("INSERT INTO likes(id_post, id_comment, id_account) VALUES(?,?,?);", id_post, 0, uuid_user)
					if err2 != nil {
						fmt.Println(err)
					}
					fmt.Println("post liked")
				}
				if temp_count_dislike > 0 {
					_, err3 := db.Exec("DELETE FROM dislikes WHERE (id_account, id_post) = (?,?);", uuid_user, id_post)
					if err3 != nil {
						fmt.Println(err)
					}
					fmt.Println("Remove Post Dislike")
				}
			} else if isComment {
				if temp_count_like > 0 {
					_, err3 := db.Exec("DELETE FROM likes WHERE (id_account, id_comment) = (?,?);", uuid_user, id_comment)
					if err3 != nil {
						fmt.Println(err)
					}
					fmt.Println(" removed comment liked by reclick")
				} else {
					_, err2 := db.Exec("INSERT INTO likes(id_post, id_comment, id_account) VALUES(?,?,?);", 0, id_comment, uuid_user)
					if err2 != nil {
						fmt.Println(err)
					}
					fmt.Println("comment liked")
				}
				if temp_count_dislike > 0 {
					_, err3 := db.Exec("DELETE FROM dislikes WHERE (id_account, id_comment) = (?,?);", uuid_user, id_comment)
					if err3 != nil {
						fmt.Println(err)
					}
					fmt.Println("Remove comment Dislike")
				}
			}
		} else if like_or_dislike == "dislike" {
			if isPost {
				if temp_count_dislike > 0 {
					_, err3 := db.Exec("DELETE FROM dislikes WHERE (id_account, id_post) = (?,?);", uuid_user, id_post)
					if err3 != nil {
						fmt.Println(err)
					}
					fmt.Println(" removed post disliked by reclick")
				} else {
					_, err2 := db.Exec("INSERT INTO dislikes(id_post, id_comment, id_account) VALUES(?,?,?);", id_post, 0, uuid_user)
					if err2 != nil {
						fmt.Println(err)
					}
					fmt.Println("post disliked")
				}
				if temp_count_like > 0 {
					_, err4 := db.Exec("DELETE FROM likes WHERE (id_account, id_post) = (?,?);", uuid_user, id_post)
					if err4 != nil {
						fmt.Println(err)
					}
					fmt.Println("Remove Post like")
				}
			} else if isComment {
				if temp_count_dislike > 0 {
					_, err3 := db.Exec("DELETE FROM dislikes WHERE (id_account, id_comment) = (?,?);", uuid_user, id_comment)
					if err3 != nil {
						fmt.Println(err)
					}
					fmt.Println(" removed comment disliked by reclick")
				} else {
					_, err2 := db.Exec("INSERT INTO dislikes(id_post, id_comment, id_account) VALUES(?,?,?);", 0, id_comment, uuid_user)
					if err2 != nil {
						fmt.Println(err)
					}
					fmt.Println("comment disliked")
				}
				if temp_count_like > 0 {
					_, err4 := db.Exec("DELETE FROM likes WHERE (id_account, id_comment) = (?,?);", uuid_user, id_comment)
					if err4 != nil {
						fmt.Println(err)
					}
					fmt.Println("Remove comment like")
				}
			}
		}
	}
}
