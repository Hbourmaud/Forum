package forum

import (
	"database/sql"
	"fmt"
)

func filters(category string, user_uuid string, created_posts string, liked_posts string) []Info {

	// variable post
	var Id_postTab []int
	var Id_accountTab []string
	var TextsTab []string
	var TitleTab []string
	var CategoryTab []string
	// var LikeTab []int
	// var DislikeTab []int
	// var One_commentTab []string
	// var One_comment_authorTab []string

	//Cela permet d'ouvrir et de fermer la database
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	var DataTab []Info

	// Je récupère les informations des posts avec la catégorie choisie par la personne
	if category == "all" {
		rows, err := db.Query("SELECT * FROM posts;", category)
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			var id_post int
			var id_account string
			var texts string
			var title string
			var category string
			var picture string

			err = rows.Scan(&id_post, &id_account, &title, &texts, &category, &picture)

			if err != nil {
				fmt.Println(err)
			}

			Id_postTab = append(Id_postTab, id_post)
			Id_accountTab = append(Id_accountTab, id_account)
			TextsTab = append(TextsTab, texts)
			TitleTab = append(TitleTab, title)
			CategoryTab = append(CategoryTab, category)
		}

		var post Info

		for i := 0; i < len(Id_postTab); i++ {
			post.Id_post = Id_postTab[i]
			post.Id_account = Id_accountTab[i]
			post.Title = TitleTab[i]
			post.Category = CategoryTab[i]
			post.Texts = TextsTab[i]

			DataTab = append(DataTab, post)
		}
		return DataTab
	} else if category == "" {
		if liked_posts == "liked_posts" {
			rows, err := db.Query("SELECT id, texts, title, category  FROM posts p JOIN likes l ON l.id_post = p.id WHERE l.id_account = (?);", user_uuid)
			if err != nil {
				fmt.Println(err)
			}
			for rows.Next() {
				var id_post int
				var id_account string
				var texts string
				var title string
				var category string

				err = rows.Scan(&id_post, &texts, &title, &category)

				if err != nil {
					fmt.Println(err)
				}

				Id_postTab = append(Id_postTab, id_post)
				Id_accountTab = append(Id_accountTab, id_account)
				TextsTab = append(TextsTab, texts)
				TitleTab = append(TitleTab, title)
				CategoryTab = append(CategoryTab, category)
			}

			var post Info

			for i := 0; i < len(Id_postTab); i++ {
				post.Id_post = Id_postTab[i]
				post.Id_account = Id_accountTab[i]
				post.Title = TitleTab[i]
				post.Category = CategoryTab[i]
				post.Texts = TextsTab[i]

				DataTab = append(DataTab, post)
			}
			return DataTab
		} else if created_posts == "created_posts" {
			rows, err := db.Query("SELECT * FROM posts WHERE id_account=(?);", user_uuid)
			if err != nil {
				fmt.Println(err)
			}
			for rows.Next() {
				var id_post int
				var id_account string
				var texts string
				var title string
				var category string
				var picture string

				err = rows.Scan(&id_post, &id_account, &title, &texts, &category, &picture)

				if err != nil {
					fmt.Println(err)
				}

				Id_postTab = append(Id_postTab, id_post)
				Id_accountTab = append(Id_accountTab, id_account)
				TextsTab = append(TextsTab, texts)
				TitleTab = append(TitleTab, title)
				CategoryTab = append(CategoryTab, category)
			}

			var post Info

			for i := 0; i < len(Id_postTab); i++ {
				post.Id_post = Id_postTab[i]
				post.Id_account = Id_accountTab[i]
				post.Title = TitleTab[i]
				post.Category = CategoryTab[i]
				post.Texts = TextsTab[i]

				DataTab = append(DataTab, post)
			}
			return DataTab
		}
	} else {
		rows, err := db.Query("SELECT * FROM posts WHERE category=(?);", category)
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			var id_post int
			var id_account string
			var texts string
			var title string
			var category string
			var picture string

			err = rows.Scan(&id_post, &id_account, &title, &texts, &category, &picture)

			if err != nil {
				fmt.Println(err)
			}

			Id_postTab = append(Id_postTab, id_post)
			Id_accountTab = append(Id_accountTab, id_account)
			TextsTab = append(TextsTab, texts)
			TitleTab = append(TitleTab, title)
			CategoryTab = append(CategoryTab, category)
		}

		var post Info

		for i := 0; i < len(Id_postTab); i++ {
			post.Id_post = Id_postTab[i]
			post.Id_account = Id_accountTab[i]
			post.Title = TitleTab[i]
			post.Category = CategoryTab[i]
			post.Texts = TextsTab[i]

			DataTab = append(DataTab, post)
		}
		return DataTab
	}
	return DataTab
}
