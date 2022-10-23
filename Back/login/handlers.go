package login

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	UserId   int    `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/login.html")
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("field-email")
	pass := r.FormValue("field-pass")
	users := UserLoginDb(email, pass)
	for _, value := range users {
		if email == value.Email && pass == value.Password {
			http.Redirect(w, r, fmt.Sprintf("/user/%d", value.UserId), http.StatusSeeOther)
		}
	}
}
