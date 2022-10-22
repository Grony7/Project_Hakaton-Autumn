package main

import (
  "html/template"
  "log"
  "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
  temp, err := template.ParseFiles("templates/index.html")
  if err != nil {
    log.Fatal(err)
  }
  temp.Execute(w, nil)
}

func handleFunc() {
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
  http.HandleFunc("/", index)
  http.ListenAndServe(":8080", nil)
}

func main() {
  handleFunc()
}
