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
<<<<<<< HEAD
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index.Index)
	rtr.HandleFunc("/login", login.Login)
	rtr.HandleFunc("/userlogin", login.UserLogin)
	rtr.HandleFunc("/registration", reg.RegPage)
	rtr.HandleFunc("/newuser", reg.Reg)
	rtr.HandleFunc("registrationerr", reg.RegErr)
	rtr.HandleFunc("/user/{id:[0-9]+}", user_page.UserPage)
	http.Handle("/", rtr)
=======
	http.HandleFunc("/", index.Index)
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/userlogin", login.Userlogin)
	http.HandleFunc("/registration", reg.RegPage)
>>>>>>> 92f15a8a728155e35d7c678b80eb2c4716ffa2b7
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
