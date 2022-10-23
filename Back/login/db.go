package login

import (
  "database/sql"
  "fmt"
)

func UserLoginDb(email string, password string) []User {
  db, err := sql.Open("mysql", "Zmeev:Sergey00616@tcp(127.0.0.1:3306)/hackaton")

  if err != nil {
    panic(err)
  }

  defer db.Close()

  data, err := db.Query(fmt.Sprintf("SELECT `user_id`, `email`, `password` FROM `users`"+
    " WHERE `email` = '%s' AND `password` = '%s'", email, password))
  if err != nil {
    panic(err)
  }

  defer data.Close()
  var userData = []User{}
  for data.Next() {
    var userEnter User
    err := data.Scan(&userEnter.UserId, &userEnter.Email, &userEnter.Password)
    if err != nil {
      panic(err)
    }
    userData = append(userData, userEnter)
  }
  return userData
}
