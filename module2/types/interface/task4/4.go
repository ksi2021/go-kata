// You can edit this code!
// Click here and start typing.
package main

import "fmt"

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
	var i Userer = &User{}
	_ = i
	fmt.Println("Success!")
}
