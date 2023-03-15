package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := sql.Open("sqlite3", "./test.db")
	// if er != nil {
	// 	panic(er)
	// }
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people(id INTEGER PRIMARY KEY, firstname TEXT, lastname Text)")
	statement.Exec()

	statement, _ = database.Prepare("INSERT INTO people(firstname, lastname) VALUES (?,?)")
	statement.Exec("testN", "testS")

	rows, _ := database.Query("SELECT * FROM people")

	var (
		id         int
		firstname  string
		secondname string
	)

	for rows.Next() {
		rows.Scan(&id, &firstname, &secondname)
		fmt.Printf("%d: %s %s \n", id, firstname, secondname)
	}
}
