package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

//type Author struct {
//	ID    uint   `gorm:"primaryKey"`
//	Name  string `gorm:"unique"`
//	Books []Book `gorm:"foreignKey:AuthorID"` // связь с таблицей книги
//}
//
//type Book struct {
//	ID       uint `gorm:"primaryKey"`
//	Title    string
//	AuthorID uint
//	Author   Author `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
//}
//
//type User struct {
//	ID    uint `gorm:"primaryKey"`
//	Name  string
//	Books []Book `gorm:"many2many:user_books"` // связь многие ко многим с таблицей
//}

type User struct {
	ID          uint   `gorm:"primaryKey"`        // ID как первичный ключ
	Name        string `fake:"{number:1,10}"`     // имя пользователя
	RentedBooks []Book `gorm:"foreignKey:UserID"` // связь с таблицей книги
}

// Book структура для таблицы книги
type Book struct {
	ID       uint   `gorm:"primaryKey"`                                                 // ID как первичный ключ
	Title    string `fake:"{sentence:3}"`                                               // название книги
	AuthorID uint   `fake:"{number:1,10}"`                                              // ID автора как внешний ключ
	Author   Author `gorm:"foreignKey:AuthorID"`                                        // связь с таблицей автор
	UserID   uint   `gorm:"default:NULL;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // ID пользователя как внешний ключ
	User     User   `gorm:"foreignKey:UserID"`                                          // связь с таблицей пользователь
}

// Author структура для таблицы автор
type Author struct {
	ID    uint   `gorm:"primaryKey" `         // ID как первичный ключ
	Name  string `fake:"{firstname}"`         // имя автора
	Books []Book `gorm:"foreignKey:AuthorID"` // связь с таблицей книги
}

func FillUsers(db *gorm.DB, count int) {
	var users []User
	var user User
	db.Find(&users)
	if len(users) < count {
		count := count - len(users)
		for i := 0; i < count; i++ {
			user = User{Name: gofakeit.FirstName() + " " + gofakeit.LastName()}
			db.Create(&user)
		}
	}

}
func FillAuthors(db *gorm.DB, count_ int) {
	var authors []Author
	var author Author
	db.Find(&authors)
	if len(authors) < count_ {
		count := count_ - len(authors)
		for i := 0; i < count; i++ {
			author = Author{Name: gofakeit.FirstName() + " " + gofakeit.LastName()}
			db.Create(&author)
		}
	}
}
func FillBooks(db *gorm.DB, count_ int) {
	var books []Book
	var book Book
	db.Find(&books)
	if len(books) < count_ {
		count := count_ - len(books)

		for i := 0; i < count; i++ {
			num := uint(gofakeit.Number(1, 10))
			book = Book{Title: gofakeit.Sentence(3), AuthorID: num}
			db.Create(&book)
		}
	}
}
