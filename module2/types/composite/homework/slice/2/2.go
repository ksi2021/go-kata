package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{
			Name: "Alice",
			Age:  21,
		},
		{
			Name: "John",
			Age:  34,
		},
		{
			Name: "Alexander",
			Age:  45,
		},
		{
			Name: "Ivan",
			Age:  13,
		},
		{
			Name: "Denis",
			Age:  44,
		},
		{
			Name: "Mary",
			Age:  26,
		},
		{
			Name: "Rose",
			Age:  41,
		},
	}

	fmt.Println("before ", users, len(users), cap(users))

	temp := []User{}
	for i := range users {
		if users[i].Age <= 40 {
			temp = append(temp, users[i])
		}
	}
	users = temp
	fmt.Println("After ", users, len(users), cap(users))

}
