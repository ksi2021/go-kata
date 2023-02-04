package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	var pop, shift User

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
	fmt.Println("Array = ", users)

	pop, users = users[0], users[1:]
	shift, users = users[len(users)-1], users[:len(users)-1]

	fmt.Println("Pop = ", pop)
	fmt.Println("Shift = ", shift)
	fmt.Println("Array = ", users)
}
