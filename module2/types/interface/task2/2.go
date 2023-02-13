// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
}

func (u *User) GetName() string {
	return u.Name
}

type Userer interface {
	GetName() string
}

func main() {
	var u Userer = &User{
		ID:   34,
		Name: "Annet",
	}
	user := u.(*User)
	testUserName(*user)
}

func testUserName(u User) {
	if u.GetName() == "Annet" {
		fmt.Println("Success!")
	}
}
