package reg

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Id       int
	UserId   int
	Name     string
	Username string
	Email    string
	Password string
	Role     string
}

func RegPage(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/register.html")
	if err != nil {
		log.Fatal(err)
	}

	temp.Execute(w, nil)
}
