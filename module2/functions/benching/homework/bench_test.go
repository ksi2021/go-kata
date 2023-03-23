// fail { не работает как надо \ переделать с нуля}
package main

/*
go test -benchmem -bench=. -v -benchtime=100x

Вроде-бы выдёт кол-во ns/op
*/

import (
	"math/rand"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

type User struct {
	ID       int64
	Name     string `fake:"{firstname}"`
	Products []Product
}

type Product struct {
	UserID int64
	Name   string `fake:"{sentence:3}"`
}

func Select(b *testing.B, users []User, products []Product) {
	// Проинициализируйте карту продуктов по айди пользователей

	for i, user := range users {
		for _, product := range products {
			if product.UserID == user.ID {
				users[i].Products = append(users[i].Products, product)
			}
		}
	}

}

func Select2(b *testing.B, users map[int]User, products []Product) {
	// Проинициализируйте карту продуктов по айди пользователей

	for _, v := range products {
		temp := users[int(v.UserID)]
		temp.Products = append(users[int(v.UserID)].Products, v)
		users[int(v.UserID)] = temp
	}

}
func BenchmarkSample100(b *testing.B) {
	users := genUsers()
	products := genProducts()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Select(b, users, products)
	}
}

func BenchmarkSample2_100(b *testing.B) {
	users := make(map[int]User)
	products := genProducts()
	temp := genUsers()
	for i := range temp {
		users[int(temp[i].ID)] = temp[i]
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Select2(b, users, products)
	}
}

func genProducts() []Product {
	products := make([]Product, 1000)
	for i, product := range products {
		_ = gofakeit.Struct(&product)
		product.UserID = int64(rand.Intn(100) + 1)
		products[i] = product
	}

	return products
}

func genUsers() []User {
	users := make([]User, 100)
	for i, user := range users {
		_ = gofakeit.Struct(&user)
		user.ID = int64(i)
		users[i] = user
	}

	return users
}
