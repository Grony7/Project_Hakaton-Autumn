package reg
<<<<<<< HEAD

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

func NewUserFunc(user_id int, name string, email string, password string, role string) {
  db, err := sql.Open("mysql", "Zmeev:Sergey00616@tcp(127.0.0.1:3306)/hackaton")
  if err != nil {
    panic(err)
  }

  defer db.Close()

  insert, err := db.Query(fmt.Sprintf("INSERT INTO `users` ("+
    "`user_id`, `name`, `email`, `password`, `role`) VALUES "+
    "('%d', '%s', '%s', '%s', '%s')", user_id, name, email, password, role))

  if err != nil {
    panic(err)
  }

  defer insert.Close()

}
=======
>>>>>>> 92f15a8a728155e35d7c678b80eb2c4716ffa2b7
