package reg

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

type NewUser struct {
	UserId   int    `json:"userId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type User struct {
	Id       int
	UserId   int
	Name     string
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

func Reg(w http.ResponseWriter, r *http.Request) {
	user_id := rand.Intn(999999)
	pass1 := r.FormValue("field-pass")
	pass2 := r.FormValue("field-pass-repeat")
	if pass1 == pass2 {
		newUser := NewUser{UserId: user_id,
			Name:     r.FormValue("field-name"),
			Email:    r.FormValue("field-email"),
			Password: r.FormValue("field-pass"),
			Role:     r.FormValue("product-group"),
		}
		NewUserFunc(newUser.UserId, newUser.Name, newUser.Email, newUser.Password, newUser.Role)
	} else {
		http.Redirect(w, r, "/reg_err", http.StatusSeeOther)
	}
}

func RegErr(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/register-error.html")
	if err != nil {
		log.Fatal(err)
	}

	temp.Execute(w, nil)
}
