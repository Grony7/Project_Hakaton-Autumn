package reg

import (
	"html/template"
	"log"
	"net/http"
)

<<<<<<< HEAD
type NewUser struct {
	UserId   int    `json:"userId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
=======
type User struct {
	Id       int
	UserId   int
	Name     string
	Username string
	Email    string
	Password string
	Role     string
>>>>>>> 92f15a8a728155e35d7c678b80eb2c4716ffa2b7
}

func RegPage(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/register.html")
	if err != nil {
		log.Fatal(err)
	}

	temp.Execute(w, nil)
}
