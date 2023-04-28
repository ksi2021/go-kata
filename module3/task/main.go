package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type DoubleLinkedList struct {
	head *Node // начальный элемент в списке
	tail *Node // последний элемент в списке
	curr *Node // текущий элемент меняется при использовании методов next, prev
	len  int   // количество элементов в списке
}

type LinkedLister interface {
	Len() int
	Current() *Node
	Next() *Node
	Prev() *Node
	LoadData(path string) error
}

// LoadData загрузка данных из подготовленного json файла
func (d *DoubleLinkedList) LoadData(path string) error {
	// отсортировать список используя самописный QuickSort

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var commits []Commit

	err = jsoniter.Unmarshal(data, &commits)

	if err != nil {
		return err
	}
	sort_commits := QuickSort(commits)
	// fmt.Println(sort_commits[0])
	// panic(nil)

	for k, v := range sort_commits {
		fmt.Println(k+1, " : ", v)
	}
	fmt.Println("------------------------------------")
	//
	if len(sort_commits) < 1 {
		return errors.New("not enought items")
	}

	var current *Node
	for _, v := range sort_commits {
		newNode := &Node{data: &Commit{Message: v.Message, UUID: v.UUID, Date: v.Date}}

		if d.head == nil {
			d.head = newNode
			d.tail = newNode
		} else {
			current = d.head
			for current.next != nil {
				current = current.next
			}
			newNode.prev = current
			current.next = newNode

			d.tail = newNode
		}

		// current = newNode

		d.len++
	}

	// d.tail = current
	d.curr = d.head
	return nil
}

// Len получение длины списка
func (d *DoubleLinkedList) Len() int {
	return d.len
}

// Current получение текущего элемента
func (d *DoubleLinkedList) Current() *Node {
	if d.curr == nil {
		d.curr = d.head
	}
	return d.curr
}

// Next получение следующего элемента
func (d *DoubleLinkedList) Next() *Node {
	if d.curr.next == nil {
		return nil
	}
	d.curr = d.curr.next
	return d.curr
}

// Prev получение предыдущего элемента
func (d *DoubleLinkedList) Prev() *Node {
	if d.curr.prev == nil {
		return nil
	}
	d.curr = d.curr.prev
	return d.curr
}

// Insert вставка элемента после n элемента
func (d *DoubleLinkedList) Insert(n int, c Commit) error {
	if n < 1 || n > d.len {
		return errors.New("index out of range")
	}

	startNode := d.head
	elementCounter := 1

	for startNode.next != nil || elementCounter <= d.len {
		if elementCounter == n {

			newNode := &Node{data: &c}

			if elementCounter == d.len {
				startNode.next = newNode
				newNode.prev = startNode
				d.len++
				return nil
			}

			afterInsertNode := startNode.next

			startNode.next = newNode
			newNode.prev = startNode
			newNode.next = afterInsertNode
			afterInsertNode.prev = newNode

			d.len++
			return nil
		}
		elementCounter++
		startNode = startNode.next
	}

	return errors.New("index out of range")
}

// Delete удаление n элемента
func (d *DoubleLinkedList) Delete(n int) error {
	panic("implement me")
}

// DeleteCurrent удаление текущего элемента
func (d *DoubleLinkedList) DeleteCurrent() error {
	if d.curr == nil {
		return errors.New("current not found")
	}
	prev, next := d.curr.prev, d.curr.next

	if prev != nil {
		if next != nil {
			prev.next = next
			next.prev = prev
			d.curr = prev
		} else {
			prev.next = nil
			d.curr = prev
		}
		d.len--
		return nil
	} else {
		if next != nil {
			next.prev = nil
			d.curr = next
			d.len--
			return nil
		} else {
			d.curr = nil
		}
	}
	return errors.New("current not found")
}

// Index получение индекса текущего элемента
func (d *DoubleLinkedList) Index() (int, error) {

	if d.curr == nil {
		return 0, errors.New("some error")
	}

	counter := 0
	start := d.head
	for start.next != nil {
		if start.data.UUID == d.curr.data.UUID {
			return counter, nil
		}
		counter++
	}

	return 0, errors.New("some error")
}

// Pop Операция Pop
func (d *DoubleLinkedList) Pop() *Node {
	last := d.tail
	prev := d.tail.prev
	if last == nil {
		return nil
	}
	if prev == nil {
		d.tail = nil
		return last
	}
	d.tail = prev
	d.tail.next = nil

	d.len--
	return last
}

// Shift операция shift
func (d *DoubleLinkedList) Shift() *Node {
	if d.head == nil {
		return nil
	}
	var oldNode *Node
	if d.head.next == nil {
		oldNode = d.head
		d.head = nil
		d.tail = nil
	} else {
		oldNode = d.head
		d.head = d.head.next
	}

	d.len--
	return oldNode

}

// SearchUUID поиск коммита по uuid
func (d *DoubleLinkedList) SearchUUID(uuID string) *Node {
	if d.len < 1 {
		return nil
	}

	SearchNode := d.head

	for SearchNode.next != nil {
		if SearchNode.data.UUID == uuID {
			return SearchNode
		}

		SearchNode = SearchNode.next
	}

	return nil

}

// Search поиск коммита по message
func (d *DoubleLinkedList) Search(message string) *Node {
	SearchNode := d.head

	for SearchNode.next != nil {
		if SearchNode.data.Message == message {
			return SearchNode
		}
		SearchNode = SearchNode.next
	}
	return nil
}

// Reverse возвращает перевернутый список
func (d *DoubleLinkedList) Reverse() *DoubleLinkedList {
	panic("implement me")
}

type Node struct {
	data *Commit
	prev *Node
	next *Node
}

type Commit struct {
	Message string    `json:"message"`
	UUID    string    `json:"uuid"`
	Date    time.Time `json:"date"`
}

func GenerateJSON() {
	// Дополнительное задание написать генератор данных
	// используя библиотеку gofakeit
}

func QuickSort(arr []Commit) []Commit {
	if len(arr) < 2 {
		return arr
	}
	max := arr[rand.Intn(len(arr))]

	l := make([]Commit, 0, len(arr))
	m := make([]Commit, 0, len(arr))
	h := make([]Commit, 0, len(arr))

	for _, v := range arr {
		switch {
		case v.Date.Before(max.Date):
			l = append(l, v)
		case v.Date.Equal(max.Date):
			m = append(m, v)
		case v.Date.After(max.Date):
			h = append(h, v)
		}
	}

	l = QuickSort(l)
	h = QuickSort(h)

	l = append(l, m...)
	l = append(l, h...)

	return l
}

func main() {
	test := &DoubleLinkedList{}

	err := test.LoadData("./test.json")
	if err != nil {
		panic(err)
	}

	// c := Commit{Message: "test message"}
	// fmt.Println(test.Insert(0, c))
	fmt.Printf("test 1 -- %+v \n", test.Current().data)
	fmt.Printf("test 2 -- %+v \n", test.Next().data)
	// fmt.Printf("test 3 -- %+v \n", test.Current().data)
	fmt.Printf("test delete -- %+v \n", test.DeleteCurrent())
	fmt.Printf("test delete -- %+v \n", test.DeleteCurrent())
	fmt.Printf("test 5 -- %+v \n", test.Current().data)
	// fmt.Printf("test 6 -- %+v \n", test.Next().data)
	// fmt.Printf("test 7 -- %+v \n", test.Next().data)
	// fmt.Printf("test 8 -- %+v \n", test.Next().data)
	// fmt.Printf("test 9 -- %+v \n", test.Next().data)
	// fmt.Printf("test 10 -- %+v \n", test.Next().data)
	fmt.Printf("test head -- %+v \n", test.tail.data)

	// fmt.Println(test.Insert(1, c))
	// fmt.Println(test.SearchUUID("6957a0ae-875b-11ed-8150-acde48001122213"))
	// fmt.Println(test.Pop().data.Message)
	// fmt.Println(test.tail.prev.data.Message)
	// fmt.Println(test.Shift().data.Message)
	// fmt.Println(test.Shift().data.Message)
	// fmt.Println(test.Shift().data.Message)
	// fmt.Println(test.Shift().data.Message)
	// fmt.Println(test.Shift().data.Message)
	// fmt.Println(test.Shift().data.Message)
	// fmt.Println(test.Shift().data.Message)
	// fmt.Println(test.Shift())
	fmt.Printf("len-- %+v \n ", test.Len())

}
