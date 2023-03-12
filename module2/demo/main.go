package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Reverser interface {
	Reverse()
}

type Sorter interface {
	Sort()
}

type SorterReverser interface {
	Reverser
	Sorter
}

type Numbers []int

func (n Numbers) Len() int {
	return len(n)
}

func (n Numbers) Less(i, j int) bool {
	return n[i] > n[j]
}

func (n Numbers) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n Numbers) Sort() { // функция сортировки массива, использовать библиотеку sort
	// sort.Ints(n)
	sort.Sort(n)
	fmt.Println(n)
}

func (n Numbers) Reverse() { // функция переворачивания массива, перевернуть используя метод a, b = b, a
	length := len(n)
	for i := 0; i < length/2; i++ {
		(n)[i], (n)[length-i-1] = (n)[length-i-1], (n)[i]
	}
	fmt.Println(n)

}

func main() {
	var numbers Numbers
	AddNumbers(&numbers) // исправить чтобы добавляло значения
	fmt.Println(numbers)
	SortReverse(numbers) // исправить чтобы сортировало и переворачивала массив

}

func AddNumbers(s *Numbers) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 11; i++ {
		*s = append(*s, rand.Intn(100))
	}
}

func SortReverse(sr SorterReverser) {
	sr.Sort()
	sr.Reverse()
}
