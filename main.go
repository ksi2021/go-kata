package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // драйвер для sqlite
)

// User - структура для представления пользователя
type Test struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	State string `db:"state"`
}

func main() {
	// Подключаемся к базе данных sqlite
	db, err := sqlx.Connect("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создаем таблицу пользователей, если ее нет
	db.MustExec(`CREATE TABLE IF NOT EXISTS tests (
		id INTEGER NOT NULL PRIMARY KEY,
		name VARCHAR(30) NOT NULL,
		state VARCHAR(30) NOT NULL
	)`)

	// Добавляем несколько пользователей в таблицу
	//tests := []Test{
	//	{State: "test1", Name: "Smith"},
	//	{State: "test2", Name: "Smith"},
	//	{State: "test3", Name: "Smith"},
	//	,
	//}
	//

	fmt.Println(List(db))
}

func Create(db *sqlx.DB, test Test) error {
	_, err := db.NamedExec(`INSERT INTO tests (name, state) VALUES (:name, :state)`, test)
	if err != nil {
		return err
	}

	return nil
}
func Delete(db *sqlx.DB, id int) error {
	if _, err := db.Exec("DELETE FROM tests WHERE id == ?;", id); err != nil {
		return err
	}
	return nil
}
func Update(db *sqlx.DB, id int, test Test) error {
	if _, err := db.Exec("UPDATE tests SET name = ?, state = ? WHERE id == ?;", test.Name, test.State, id); err != nil {
		return err
	}
	return nil
}
func List(db *sqlx.DB) ([]Test, error) {
	tests := []Test{}

	err := db.Select(&tests, "SELECT * FROM tests")
	if err != nil {
		return []Test{}, err
	}

	return tests, nil
}
