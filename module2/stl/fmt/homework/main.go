package main

import (
	"fmt"
)

func generateSelfStory(name string, age int, money float64) string {
	return fmt.Sprintf("Hello! My name is %s. I`m %d y.o. And I also have $%.2f in my wallet fight now.", name, age, money)
}

func main() {
	fmt.Println(generateSelfStory("Vlad", 25, 10.0000025))
}
