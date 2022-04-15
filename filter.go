package forum

import (
	"database/sql"
	"fmt"
)

func filters(category string) []Info {

	// variable post
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

	// // Je récupère les informations des posts avec la catégorie choisie par la personne
	rows, err := db.Query("SELECT * FROM posts WHERE category=(?);", category)
	if err != nil {
		fmt.Println(err)
	}

	var DataTab []Info
	fmt.Println(DataTab)
	for rows.Next() {
		var id_post int
		var id_account string
		var picture_text string
		var title string
		var category string

		err = rows.Scan(&id_post, &id_account, &picture_text, &title, &category)

		if err != nil {
			fmt.Println(err)
		}

		Id_postTab = append(Id_postTab, id_post)
		Id_accountTab = append(Id_accountTab, id_account)
		Picture_textTab = append(Picture_textTab, picture_text)
		TitleTab = append(TitleTab, title)
		CategoryTab = append(CategoryTab, category)
	}

	fmt.Println(CategoryTab)

	var post Info

	for i := 0; i < len(Id_postTab); i++ {
		post.Id_post = Id_postTab[i]
		post.Id_account = Id_accountTab[i]
		post.Title = TitleTab[i]
		post.Category = CategoryTab[i]
		post.Picture_text = Picture_textTab[i]

		DataTab = append(DataTab, post)
	}
	fmt.Println(DataTab)
	return DataTab
}
