package forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type NewData struct {
	DataTab            []Info
	DataTabCreate      []Info
	DataTabPostLike    []Info
	DataTabPostDislike []Info
	DataTabComment     []Info
}

type ActivityComment struct {
	OneCommentTab []string
	DataComment   []Info
}

func Activity(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/activity.html"))

	ck_uuid_user, err := r.Cookie("uuid_hash")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		fmt.Println(err)
		return
	}
	uuid_user := ck_uuid_user.Value
	r.ParseForm()

	/* postes qu'il a créé */

	// variable post
	var Id_postTab []int
	var Id_accountTab []string
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
	var DataTabCreate []Info
	// var DataTabComment []Info

	// Je récupère les informations
	rows, err := db.Query("SELECT * FROM posts WHERE id_account=(?);", uuid_user)
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

		err = rows.Scan(&id_post, &id_account, &texts, &title, &category, &picture)

		if err != nil {
			fmt.Println(err)
		}

		Id_postTab = append(Id_postTab, id_post)
		Id_accountTab = append(Id_accountTab, id_account)
		TextsTab = append(TextsTab, texts)
		TitleTab = append(TitleTab, title)
		CategoryTab = append(CategoryTab, category)
		PictureTab = append(PictureTab, picture)
	}
	defer rows.Close()
	var post Info

	for i := 0; i < len(Id_postTab); i++ {
		post.Id_post = Id_postTab[i]
		post.Id_account = idAccount_to_username(Id_accountTab[i])
		post.Title = TitleTab[i]
		post.Category = CategoryTab[i]
		post.Texts = TextsTab[i]
		post.Picture = PictureTab[i]

		DataTabCreate = append(DataTabCreate, post)
	}

	// variable post
	var Id_postTab1 []int
	var Id_accountTab1 []string
	var TextsTab1 []string
	var TitleTab1 []string
	var CategoryTab1 []string
	var PictureTab1 []string

	var DataTabPostLike []Info
	rows1, err := db.Query("SELECT id,p.id_account, texts, title, category, picture  FROM posts p JOIN likes l ON l.id_post = p.id WHERE l.id_account = (?);", uuid_user)
	if err != nil {
		fmt.Println(err)
	}
	for rows1.Next() {
		var id_post int
		var id_account string
		var texts string
		var title string
		var category string
		var picture string

		err = rows1.Scan(&id_post, &id_account, &texts, &title, &category, &picture)

		if err != nil {
			fmt.Println(err)
		}

		Id_postTab1 = append(Id_postTab1, id_post)
		Id_accountTab1 = append(Id_accountTab1, id_account)
		TextsTab1 = append(TextsTab1, texts)
		TitleTab1 = append(TitleTab1, title)
		CategoryTab1 = append(CategoryTab1, category)
		PictureTab1 = append(PictureTab1, picture)
	}
	defer rows1.Close()
	var post1 Info

	for i := 0; i < len(Id_postTab1); i++ {
		post1.Id_post = Id_postTab1[i]
		post1.Id_account = idAccount_to_username(Id_accountTab1[i])
		post1.Title = TitleTab1[i]
		post1.Category = CategoryTab1[i]
		post1.Texts = TextsTab1[i]
		post1.Picture = PictureTab1[i]

		DataTabPostLike = append(DataTabPostLike, post1)
	}

	// variable post
	var Id_postTab2 []int
	var Id_accountTab2 []string
	var TextsTab2 []string
	var TitleTab2 []string
	var CategoryTab2 []string
	var PictureTab2 []string

	var DataTabPostDislike []Info
	rows2, err := db.Query("SELECT id,p.id_account, texts, title, category, picture  FROM posts p JOIN dislikes l ON l.id_post = p.id WHERE l.id_account = (?);", uuid_user)
	if err != nil {
		fmt.Println(err)
	}
	for rows2.Next() {
		var id_post int
		var id_account string
		var texts string
		var title string
		var category string
		var picture string

		err = rows2.Scan(&id_post, &id_account, &texts, &title, &category, &picture)

		if err != nil {
			fmt.Println(err)
		}

		Id_postTab2 = append(Id_postTab2, id_post)
		Id_accountTab2 = append(Id_accountTab2, id_account)
		TextsTab2 = append(TextsTab2, texts)
		TitleTab2 = append(TitleTab2, title)
		CategoryTab2 = append(CategoryTab2, category)
		PictureTab2 = append(PictureTab2, picture)
	}
	defer rows2.Close()
	var post2 Info

	for i := 0; i < len(Id_postTab2); i++ {
		post2.Id_post = Id_postTab2[i]
		post2.Id_account = idAccount_to_username(Id_accountTab2[i])
		post2.Title = TitleTab2[i]
		post2.Category = CategoryTab2[i]
		post2.Texts = TextsTab2[i]
		post2.Picture = PictureTab2[i]

		DataTabPostDislike = append(DataTabPostDislike, post2)
	}
	var DataTabComment []CommentStruct
	var DataTabCommentPost []Info

	rows3, err := db.Query("SELECT id,id_post,comment FROM comments WHERE id_account = (?);", uuid_user)
	if err != nil {
		fmt.Println(err)
	}
	defer rows3.Close()
	for rows3.Next() {
		var comment CommentStruct
		var id_post int
		var id_comment int
		var comment_content string
		err = rows3.Scan(&id_comment, &id_post, &comment_content)
		if err != nil {
			fmt.Println(err)
		}
		comment.Comment = comment_content
		comment.Id_post = id_post
		comment.Id_comment = id_comment
		DataTabComment = append(DataTabComment, comment)
	}
	for _, comm := range DataTabComment {
		rows4, err := db.Query("SELECT * FROM posts WHERE id = (?);", comm.Id_post)
		if err != nil {
			fmt.Println(err)
		}
		defer rows4.Close()
		for rows4.Next() {
			var id_post int
			var id_account string
			var texts string
			var title string
			var category string
			var picture string
			var post Info
			err = rows4.Scan(&id_post, &id_account, &title, &texts, &category, &picture)
			if err != nil {
				fmt.Println(err)
			}
			post.Id_post = id_post
			post.Id_account = idAccount_to_username(id_account)
			post.Title = title
			post.Texts = texts
			post.Category = category
			post.Picture = picture
			DataTabCommentPost = append(DataTabCommentPost, post)
		}
	}
	DataTabCreate = addDetailsPost(DataTabCreate)
	DataTabPostLike = addDetailsPost(DataTabPostLike)
	DataTabPostDislike = addDetailsPost(DataTabPostDislike)
	DataTabCommentPost = addDetailsPost(DataTabCommentPost)
	for index := range DataTabComment {
		DataTabCommentPost[index].One_comment = DataTabComment[index].Comment
		DataTabCommentPost[index].One_comment_id = strconv.Itoa(DataTabComment[index].Id_comment)
	}

	data := &NewData{
		DataTabCreate:      DataTabCreate,
		DataTabPostLike:    DataTabPostLike,
		DataTabPostDislike: DataTabPostDislike,
		DataTabComment:     DataTabCommentPost,
	}
	err2 := tmpl.Execute(w, data)
	if err2 != nil {
		fmt.Println(err2)
	}
}
