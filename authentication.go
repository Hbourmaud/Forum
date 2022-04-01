package forum

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type DataAuthentication struct {
}

func AuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/authentication.html"))
	login := false
	data := &DataAuthentication{}
	switch r.Method {
	case "GET":

	case "POST":
		username := r.FormValue("username")
		email := r.FormValue("email")
		passwd := r.FormValue("passwd")
		email_taken := false
		if username == "" {
			login = true
		}
		db, err := sql.Open("sqlite3", "./forum.db")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		//test := bcrypt.CompareHashAndPassword(bytes, []byte("password"))
		if !login {
			rows, err := db.Query("SELECT count(*) FROM authentication WHERE email=(?);", email)
			if err != nil {
				fmt.Println(err)
			}
			defer rows.Close()
			for rows.Next() {
				var email_nb int
				err = rows.Scan(&email_nb)
				if email_nb > 0 {
					email_taken = true
				}
				if err != nil {
					fmt.Println(err)
				}
			}
			fmt.Println("email", email_taken)
			if email_taken {
				fmt.Println("This email is already taken. Choose another one!")
			} else {
				u1 := (uuid.NewV1()).String()
				crypt_passwd, err := bcrypt.GenerateFromPassword([]byte(passwd), 14)
				if err != nil {
					fmt.Println(err)
				}
				_, err = db.Exec("INSERT INTO authentication(UUID, username, email, password) VALUES(?,?,?,?);", u1, username, email, crypt_passwd)
				if err != nil {
					fmt.Println(err)
				}
			}
		} else {
			rows, err := db.Query("SELECT * FROM authentication WHERE email=(?);", email)
			if err != nil {
				fmt.Println(err)
			}
			defer rows.Close()
			for rows.Next() {
				var email_check string
				var passw_check string
				err = rows.Scan(&email_check, &passw_check)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	tmpl.Execute(w, data)
}
