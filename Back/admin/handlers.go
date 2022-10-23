package admin

import (
  "html/template"
  "log"
  "net/http"
)

type Post struct {
  Id    int    `json:"id"`
  Title string `json:"title"`
  Text  string `json:"text"`
  Desc  string `json:"desc"`
  Img   string `json:"img"`
}

func AdminPage(w http.ResponseWriter, r *http.Request) {
  temp, err := template.ParseFiles("./templates/admin.html")
  if err != nil {
    log.Fatal(err)
  }

  temp.Execute(w, nil)
}

func NewPost(w http.ResponseWriter, r *http.Request) {
  post := Post{
    Title: r.FormValue("post-title"),
    Text:  r.FormValue("post-text"),
    Img:   r.FormValue("post-img"),
    Desc:  r.FormValue("post-desc"),
  }

}
