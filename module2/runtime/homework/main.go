// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("i can manage")

	go func() {
		fmt.Println("goroutines in Golang!")
	}()
	runtime.Gosched()
	fmt.Println("and its awesome!")

}
