package admin

//
//import (
//_ "github.com/go-sql-driver/mysql"
//)
//
//func NewPost(title string, text string, desc string, img string) {
// db, err := sql.Open("mysql", "Zmeev:Sergey00616@tcp(127.0.0.1:3306)/hackaton")
// if err != nil {
//   panic(err)
// }
//
// defer db.Close()
//
// insert, err := db.Query(fmt.Sprintf("INSERT INTO `posts` ("+
//   "`title`, `text`, `description`, `img`) VALUES (`%s`, `%s`, `%s`, `%s`", title, text, desc, img))
// if err != nil {
//   panic(err)
// }
// defer insert.Close()
//
//}
