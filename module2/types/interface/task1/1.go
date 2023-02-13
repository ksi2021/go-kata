// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

type MyInterface interface{}

func main() {
	var n *int
	fmt.Println(n == nil)
	fmt.Println(test(n))
}

func test(r interface{}) string {
	switch r.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	case bool:
		return "bool"
	case []int:
		return "[]int"
	/*
		... etc
	*/
	default:
		return "Success!"
	}

}
