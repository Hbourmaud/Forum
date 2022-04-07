package forum

import (
	"fmt"
	"html/template"
	"net/http"
)

type Data struct {
	Username  string
	UUID_hash string
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	data := &Data{
		Username:  "Not login",
		UUID_hash: "",
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
		data = &Data{
			Username:  ck_user.Value,
			UUID_hash: ck_uuid.Value,
		}
	}
	tmpl.Execute(w, data)
}
