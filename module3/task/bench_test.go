// fail { не работает как надо \ переделать с нуля}
package main

/*
go test -benchmem -bench=. -v -benchtime=100x

Вроде-бы выдёт кол-во ns/op
*/

import (
	"testing"
)

func BenchmarkLoadData100(b *testing.B) {
	test := &DoubleLinkedList{}
	for i := 0; i < b.N; i++ {
		err := test.LoadData("./test.json")
		_ = err
	}
}

func BenchmarkReverseList100(b *testing.B) {
	test := &DoubleLinkedList{}
	err := test.LoadData("./test.json")
	_ = err
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reverse := test.Reverse()
		_ = reverse
	}
}

func BenchmarkPrev100(b *testing.B) {
	test := &DoubleLinkedList{}
	err := test.LoadData("./test.json")
	_ = err
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		temp := test.Prev()
		_ = temp
	}
}

func BenchmarkPop100(b *testing.B) {
	test := &DoubleLinkedList{}
	err := test.LoadData("./test.json")
	_ = err
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		temp := test.Pop()
		_ = temp
	}
}

func BenchmarkShift100(b *testing.B) {
	test := &DoubleLinkedList{}
	err := test.LoadData("./test.json")
	_ = err
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		temp := test.Shift()
		_ = temp
	}
}

func BenchmarkInsert100(b *testing.B) {
	test := &DoubleLinkedList{}
	err := test.LoadData("./test.json")
	_ = err
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		commit := Commit{}
		temp := test.Insert(11, commit)
		_ = temp
	}
}

func BenchmarkDelete100(b *testing.B) {
	test := &DoubleLinkedList{}
	err := test.LoadData("./test.json")
	_ = err
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		temp := test.Delete(11)
		_ = temp
	}
}

func BenchmarkSearch100(b *testing.B) {
	test := &DoubleLinkedList{}
	err := test.LoadData("./test.json")
	_ = err
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		temp := test.Search("Try to bypass the PCI firewall, maybe it will calculate the optical system!")
		_ = temp
	}
}

func BenchmarkDeleteCurrent100(b *testing.B) {
	test := &DoubleLinkedList{}
	err := test.LoadData("./test.json")
	_ = err
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		temp := test.DeleteCurrent()
		_ = temp
	}
}
