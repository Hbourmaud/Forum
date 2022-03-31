package forum

import (
	"html/template"
	"net/http"
)

type Data struct {
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	data := &Data{}

	tmpl.Execute(w, data)
}
