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
	Role     string `json:"role"`
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
	fmt.Println(users)
	if len(users) != 0 {
		for _, value := range users {
			if email == value.Email && pass == value.Password {
				if value.Role == "student" || value.Role == "common-user" {
					http.Redirect(w, r, fmt.Sprintf("/student/%d", value.UserId), http.StatusSeeOther)
				} else if value.Role == "professor" {
					http.Redirect(w, r, fmt.Sprintf("/prof/%d", value.UserId), http.StatusSeeOther)
				}
			}
		}
	} else {
		http.Redirect(w, r, "/log_err", http.StatusSeeOther)
	}
}

func LogErr(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/reg-error.html")
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}
