package main

import (
	"./index"
	"./login"
	"./reg"
	"./user_page"
	"github.com/gorilla/mux"
	"net/http"
)

func handleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index.Index)
	rtr.HandleFunc("/login", login.Login)
	rtr.HandleFunc("/userlogin", login.UserLogin)
	rtr.HandleFunc("/registration", reg.RegPage)
	rtr.HandleFunc("/newuser", reg.Reg)
	rtr.HandleFunc("registrationerr", reg.RegErr)
	rtr.HandleFunc("/user/{id:[0-9]+}", user_page.UserPage)
	http.Handle("/", rtr)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
