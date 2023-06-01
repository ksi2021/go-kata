package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	value int
	prev  *Node
	next  *Node
}
type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func (list *DoubleLinkedList) InsertFirst(value int) {
	newNode := &Node{value, nil, nil}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}
}

func (list *DoubleLinkedList) InsertLast(value int) {
	newNode := &Node{value, nil, nil}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
}
func (list *DoubleLinkedList) DeleteFirst() {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}
	list.head = list.head.next
	if list.head == nil {
		list.tail = nil
	} else {
		list.head.prev = nil
	}
}
func (list *DoubleLinkedList) DeleteLast() {
	if list.tail == nil {
		fmt.Println("List is empty")
		return
	}
	list.tail = list.tail.prev
	if list.tail == nil {
		list.head = nil
	} else {
		list.tail.next = nil
	}
}
func (list *DoubleLinkedList) Delete(value int) {
	node := list.Search(value)
	if node == nil {
		fmt.Println("Value not found in the list")
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		list.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		list.tail = node.prev
	}
}
func (list *DoubleLinkedList) Search(value int) *Node {
	node := list.head
	for node != nil {
		if node.value == value {
			return node
		}
		node = node.next
	}
	return nil
}

func getStringRepresentation(node *Node) string {
	result := ""
	for node != nil {
		result += strconv.Itoa(node.value) + " <=> "
		node = node.next
	}
	result += "nil"
	return result
}

func (list *DoubleLinkedList) Display() {
	node := list.head
	result := getStringRepresentation(node)
	fmt.Println(result)
}
