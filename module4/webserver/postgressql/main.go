package main

import (
	"log"

	//"github.com/go-gormigrate/gormigrate/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct {
	ID    int `gorm:"primary_key"`
	Name  string
	Email string
}

func main() {
	// Открываем соединение с БД с помощью gorm
	db, err := gorm.Open("postgres", "user=postgres password=pass dbname=kata sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if !db.HasTable(User{}) {
		db.CreateTable(User{}, "Users")
	}
	//db.Create(&User{ID: 4, Name: "test", Email: "email"})
	// Создаем экземпляр миграции с указанием таблицы для хранения истории миграций
	//m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
	//	// Описываем миграцию на основе структуры User
	//	{
	//		ID: "202111021234",
	//		Migrate: func(tx *gorm.DB) error {
	//			// Создаем таблицу users в БД, если ее нет
	//			return tx.AutoMigrate(&User{}).Error
	//		},
	//		Rollback: func(tx *gorm.DB) error {
	//			// Удаляем таблицу users из БД, если нужно откатить миграцию
	//			return tx.DropTable("users").Error
	//		},
	//	},
	//})

	// Применяем все доступные миграции к БД
	//err = m.Migrate()
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Выполняем запрос и получаем результат в виде []*User
	var users []*User
	err = db.Find(&users).Error
	if err != nil {
		log.Fatal(err)
	}

	// Выводим данные пользователей на экран
	for _, user := range users {
		log.Printf("%d %s %s\n", user.ID, user.Name, user.Email)
	}
}
