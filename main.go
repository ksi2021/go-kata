// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// получаем 2 слайса с канала
	ch := NewChanWithNumbers(1597)
	arr1, arr2 := ReadFromChannel(ch)
	// вывести 2 слайса
	fmt.Println("arr1 = ", arr1, "\n arr2 =", arr2)
	// отсортировать слайсы бабл сортом
	BubbleSort(arr1)
	BubbleSort(arr2)
	// смержить отсортированные слайсы в один отсортированный
	arr := MergeSlices(arr1, arr2)
	// выводим информацию is sorted и сам слайс
	fmt.Println(IsSorted(arr))
	fmt.Println(arr)
}

func BubbleSort(arr []int) {
	for i := range arr {
		for j := range arr {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func MergeSlices(a []int, b []int) []int {
	// смержить 2 слайса сохраняя сортировку

	newArr := make([]int, len(a)+len(b))

	copy(newArr[:len(a)], a)
	copy(newArr[len(a)+1:], b)

	BubbleSort(newArr)

	return newArr
}

func IsSorted(arr []int) bool {

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}

	return true
	// проверка что слайс отсортирован
}

func ReadFromChannel(ch chan int) ([]int, []int) {
	slice1 := make([]int, 610)
	slice2 := make([]int, 987)

	counter := 0
	// записываем значения из каналов в 2 слайса длиной 610 и 987
	for i := range ch {
		slice1[counter] = i
		counter++

		if counter >= 610-1 {
			break
		}
	}
	counter = 0
	for i := range ch {
		slice2[counter] = i
		counter++

		if counter >= 987-1 {
			break
		}
	}

	return slice2, slice1

}

func NewChanWithNumbers(capacity int) chan int {
	ch := make(chan int, capacity)

	go func() {
		// записываем в канал значения до тех пор, пока буффер не заполнится
		// Если в буфере канала есть место, отправляем новое значение
		for i := 0; i < capacity; i++ {
			ch <- rand.Intn(150)
		}

	}()

	// Возвращаем буферизированный канал
	return ch

}
