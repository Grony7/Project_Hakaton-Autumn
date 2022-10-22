package login

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/login.html")
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

func Userlogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("field-email")
	pass := r.FormValue("field-pass")

	fmt.Println(email)
	fmt.Println(pass)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
