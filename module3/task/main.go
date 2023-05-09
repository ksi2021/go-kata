package main

import (
	"errors"
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

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var commits []Commit

	err = jsoniter.Unmarshal(data, &commits)

	if err != nil {
		return err
	}

	// сортировка списка используя самописный QuickSort
	sort_commits := QuickSort(commits[:150])

	// for k, v := range sort_commits[:3] {
	// 	fmt.Println(k+1, " : ", v)
	// }
	// fmt.Println("------------------------------------")
	//
	if len(sort_commits) < 1 {
		return errors.New("not enought items")
	}

	var current *Node
	var newNode *Node = nil
	for _, v := range sort_commits {
		newNode = &Node{data: &Commit{Message: v.Message, UUID: v.UUID, Date: v.Date}}

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
func (d *DoubleLinkedList) Insert(n int, c Commit) bool {
	if n < 1 || n > d.len {
		return false
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
				return true
			}

			afterInsertNode := startNode.next

			startNode.next = newNode
			newNode.prev = startNode
			newNode.next = afterInsertNode
			afterInsertNode.prev = newNode

			d.len++
			return true
		}
		elementCounter++
		startNode = startNode.next
	}

	return true
}

// Delete удаление n элемента
func (d *DoubleLinkedList) Delete(n int) bool {
	temp := d.head
	counter := 0

	if d.len-1 == n {
		if d.tail.prev != nil {

			if d.tail == d.curr {
				d.curr = d.tail.prev
			}
			d.tail = d.tail.prev
			d.tail.next = nil

		} else {
			d.tail = nil
			d.curr = nil
			d.head = nil
		}

		d.len--
		return true
	}

	for temp.next != nil {
		if counter == n {
			next := temp.next
			prev := temp.prev
			if d.curr == temp {
				d.curr = next
			}
			if prev == nil {
				next.prev = nil
				d.head = next
			} else {
				next.prev = prev
				prev.next = next
			}

			d.len--

			return true
		}
		counter++
		temp = temp.next
	}

	if counter < n {
		return false
	}

	return false
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
			d.curr = next
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
			d.head = next
			d.len--
			return nil
		} else {
			d.head = nil
			d.tail = nil
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

	if d.len < 1 {
		return nil
	}

	del := d.tail

	if del.prev == nil {
		d.head = nil
		d.tail = nil
		d.curr = nil
	} else {
		d.tail = del.prev
		d.tail.next = nil
		if d.curr == del {
			d.curr = d.tail
		}
	}

	d.len--
	return del
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

	for SearchNode != nil {
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
	var temp *Node
	current := d.head

	for current != nil {
		temp = current.prev
		current.prev = current.next
		current.next = temp
		current = current.prev
	}

	if temp != nil {
		d.head = temp.prev
	}

	d.curr = d.head
	return d
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

	// test.Reverse()
	// fmt.Printf("el f = %+v \n ", test.Current())
	// fmt.Printf("el f = %+v \n ", test.head)
	// // fmt.Printf("el f = %+v \n ", test.Next())
	// // fmt.Printf("el f = %+v \n ", test.Next())

	// // fmt.Printf("el reverse = %+v \n ", test.Reverse())

	// // fmt.Printf("el 1 = %+v \n ", test.head.data.Message)
	// // fmt.Printf("el 2 = %+v \n ", test.head.next.data.Message)
	// // fmt.Printf("el 3 = %+v \n ", test.head.next.next.data.Message)

	// // fmt.Printf("el f = %+v \n ", test.Index(0))

	// // fmt.Printf("el reverse = %+v \n ", test.tail.data.Message)
	// // fmt.Printf("el reverse = %+v \n ", test.tail.prev.data.Message)
	// // fmt.Printf("el reverse = %+v \n ", test.tail.prev.prev.data.Message)
	// // temp := test.head
	// // i := 1
	// // for temp != nil {
	// // 	fmt.Printf("el № %d -- %+v \n", i, temp.data.Message)
	// // 	i++
	// // 	temp = temp.next
	// // }

	// // fmt.Printf("el reverse = %+v \n ", test.Reverse())

	// // temp := test.head
	// // i := 1
	// // for temp != nil {
	// // 	fmt.Printf("el № %d -- %+v \n", i, temp.data.Message)
	// // 	i++
	// // 	temp = temp.next
	// // }
	// // fmt.Printf("el curr = %+v \n ", test.Current().data.Message)
	// // fmt.Printf("el curr = %+v \n ", test.tail)
	// fmt.Printf("len-- %+v \n ", test.Len())

}
