package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{
			Name: "Eva",
			Age:  13,
		},
		{
			Name: "Victor",
			Age:  28,
		},
		{
			Name: "Dex",
			Age:  34,
		},
		{
			Name: "Billy",
			Age:  21,
		},
		{
			Name: "Foster",
			Age:  29,
		},
	}

	subUsers := make([]User, len(users[2:]))

	copy(subUsers, users[2:])
	editSecondSlice(subUsers)

	fmt.Println(users)
}

func editSecondSlice(users []User) {
	for i := range users {
		users[i].Name = "unknown"
	}
}
