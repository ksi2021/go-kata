package main

import (
	"fmt"
)

func main() {
	// 1 вариант
	{
		s := []int{1, 2, 3}
		Append(&s)
		fmt.Println(s, len(s), cap(s))
	}

	// 2 вариант
	{
		s := []int{1, 2, 3}
		s = Append2(s)
		fmt.Println(s, len(s), cap(s))
	}

}

func Append(s *[]int) {
	*s = append(*s, 4)
}

func Append2(s []int) []int {
	s = append(s, 4)
	return s
}
