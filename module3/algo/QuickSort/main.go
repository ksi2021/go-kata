package main

import (
	"fmt"
	"math/rand"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	max := arr[rand.Intn(len(arr))]

	l := make([]int, 0, len(arr))
	m := make([]int, 0, len(arr))
	h := make([]int, 0, len(arr))

	for _, v := range arr {
		switch {
		case v < max:
			l = append(l, v)
		case v == max:
			m = append(m, v)
		case v > max:
			h = append(h, v)
		}
	}

	l = quickSort(l)
	h = quickSort(h)

	l = append(l, m...)
	l = append(l, h...)

	return l
}

func main() {
	e := []int{2, 8, 56, 1, 13, 22, 24}
	test := quickSort(e)

	fmt.Println(e, test)
}
