package user_page

import (
	"html/template"
	"log"
	"net/http"
)

func UserSPage(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/index-s.html")
	if err != nil {
		log.Fatal(err)
	}

	temp.Execute(w, nil)
}

func UserPPage(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/index-teacher.html")
	if err != nil {
		log.Fatal(err)
	}

	temp.Execute(w, nil)
}
