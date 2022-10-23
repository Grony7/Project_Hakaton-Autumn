package main

import (
  "./admin"
  "./index"
  "./login"
  "./reg"
  "./user_page"
  "github.com/gorilla/mux"
  "net/http"
)

func handleFunc() {

  rtr := mux.NewRouter()
  rtr.HandleFunc("/", index.Index)
  rtr.HandleFunc("/login", login.Login)
  rtr.HandleFunc("/userlogin", login.UserLogin)
  rtr.HandleFunc("/registration", reg.RegPage)
  rtr.HandleFunc("/newuser", reg.Reg)
  rtr.HandleFunc("/reg_err", reg.RegErr)
  rtr.HandleFunc("/log_err", login.LogErr)
  rtr.HandleFunc("/student/{id:[0-9]+}", user_page.UserSPage)
  rtr.HandleFunc("/prof/{id:[0-9]+}", user_page.UserPPage)
  rtr.HandleFunc("/admin", admin.AdminPage)

  http.Handle("/", rtr)
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
  http.ListenAndServe(":8080", nil)
}

func main() {
  handleFunc()
}
