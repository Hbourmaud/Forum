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
	UUID_hash string
	Username  string
}

func AuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/authentication.html"))
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case "POST":
		username := r.FormValue("username")
		email := r.FormValue("email")
		passwd := r.FormValue("passwd")
		email_taken := false
		db, err := sql.Open("sqlite3", "./forum.db")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		rows, err2 := db.Query("SELECT count(*) FROM authentication WHERE email=(?);", email)
		if err2 != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		for rows.Next() {
			var email_nb int
			err3 := rows.Scan(&email_nb)
			if email_nb > 0 {
				email_taken = true
			}
			if err3 != nil {
				fmt.Println(err)
			}
		}
		fmt.Println("email", email_taken)
		if email_taken {
			fmt.Println("This email is already taken. Choose another one!")
		} else {
			uuid_crypted, err4 := bcrypt.GenerateFromPassword([]byte((uuid.NewV1()).String()), 14)
			if err4 != nil {
				fmt.Println(err)
			}
			crypt_passwd, err5 := bcrypt.GenerateFromPassword([]byte(passwd), 14)
			if err5 != nil {
				fmt.Println(err)
			}

			_, err6 := db.Exec("INSERT INTO authentication(UUID, username, email, password) VALUES(?,?,?,?);", uuid_crypted, username, email, crypt_passwd)
			if err6 != nil {
				fmt.Println(err)
			}
			Setter_Cookie(w, r, username, string(uuid_crypted))
		}
	}
	data := ""
	tmpl.Execute(w, data)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/login.html"))
	username := r.FormValue("username")
	email := r.FormValue("email")
	passwd := r.FormValue("passwd")
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var passw_check []byte
	var uid_hash string
	rows, err2 := db.Query("SELECT * FROM authentication WHERE email=(?);", email)
	if err2 != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var email_check string
		err3 := rows.Scan(&uid_hash, &username, &email_check, &passw_check)
		if err3 != nil {
			fmt.Println(err)
		}
	}
	correct_passwd := bcrypt.CompareHashAndPassword(passw_check, []byte(passwd))
	if correct_passwd == nil {
		Setter_Cookie(w, r, username, uid_hash)
	} else {
		fmt.Println("Wrong Password")
	}
	data := ""
	tmpl.Execute(w, data)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	delete_cookie_usr := &http.Cookie{
		Name:   "username",
		Value:  "",
		MaxAge: -1,
	}
	delete_cookie_uuid := &http.Cookie{
		Name:   "uuid_hash",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, delete_cookie_usr)
	http.SetCookie(w, delete_cookie_uuid)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Setter_Cookie(w http.ResponseWriter, r *http.Request, username string, uid_hash string) {
	ck_user := http.Cookie{
		Name: "username",
	}
	ck_uuid := http.Cookie{
		Name: "uuid_hash",
	}
	ck_user.Value = username
	ck_uuid.Value = uid_hash
	http.SetCookie(w, &ck_user)
	http.SetCookie(w, &ck_uuid)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
