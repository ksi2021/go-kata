package main

import (
	"fmt"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
)

func generateSelfStory(p Person) string {
	return fmt.Sprintf("Hello! My name is %s. I`m %d y.o. And I also have $%.2f in my wallet right now.", p.Name, p.Age, p.Money)
}

type Person struct {
	Name  string  `fake:"{firstname}"`
	Age   int     `fake:"{number:15,84}"`
	Money float64 `fake:"{float64range:0,3000}"`
}

func main() {
	var p Person = Person{}
	count := 100 + rand.Intn(1000-100+1)

	for i := 0; i < count; i++ {
		gofakeit.Struct(&p)
		fmt.Println(generateSelfStory(p))
	}
}
