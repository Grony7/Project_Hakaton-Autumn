package index

import (
  "html/template"
  "log"
  "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
  temp, err := template.ParseFiles("./templates/index.html")
  if err != nil {
    log.Fatal(err)
  }
  temp.Execute(w, nil)
}
