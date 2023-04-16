package main

import (
	"errors"
	"fmt"
	"time"
)

type Post struct {
	body        string
	publishDate int64 // Unix timestamp
	next        *Post
}

type Feed struct {
	length int // we'll use it later
	start  *Post
	end    *Post
}

// Реализуйте метод Append для добавления нового поста в конец потока
func (f *Feed) Append(newPost *Post) {
	// ваш код здесь
	if f.length == 0 {
		f.start = newPost
		f.end = newPost
	} else {
		temp := f.end
		temp.next = newPost
		f.end = newPost

	}
	f.length++
}

// Реализуйте метод Remove для удаления поста по дате публикации из потока
func (f *Feed) Remove(publishDate int64) {
	// ваш код здесь

	if f.length == 0 {
		panic(errors.New("zero length"))
	}
	var prev *Post
	current := f.start

	for current.publishDate != publishDate {

		if current.next == nil {
			panic(errors.New("post not found"))
		}
		prev = current
		current = current.next
	}
	// если удаляем 1 элемент
	if current.publishDate == f.start.publishDate {
		f.start = current.next

	} else /* все остальное */ {
		temp := current.next
		prev.next = temp
	}

	f.length--
}

// Реализуйте метод Insert для вставки нового поста в поток в соответствии с его датой публикации
func (f *Feed) Insert(newPost *Post) {
	// ваш код здесь
	if f.length == 0 {
		f.start = newPost
		f.end = newPost
	} else {
		current := f.start
		prev := current
		for i := 0; i < f.length; i++ {

			if current.next == nil {
				temp := f.end
				temp.next = newPost
				f.end = newPost
				break
			}

			if current.publishDate > newPost.publishDate {
				prev.next = newPost
				temp := prev.next
				temp.next = current
				break
			}
			prev = current
			current = current.next
		}

	}
	f.length++
}

// Реализуйте метод Inspect для вывода информации о потоке и его постах
func (f *Feed) Inspect() {
	// ваш код здесь
	fmt.Printf("Feed length: %d \n", f.length)
	current := f.start

	for i := 0; i < f.length; i++ {
		fmt.Printf("Item: %d - %v \n", i, current)
		current = current.next
	}
}

func main() {
	rightNow := time.Now().Unix()
	f := &Feed{}
	p1 := &Post{
		body:        "Lorem ipsum",
		publishDate: rightNow,
	}
	p2 := &Post{
		body:        "Dolor sit amet",
		publishDate: rightNow + 10,
	}
	p3 := &Post{
		body:        "consectetur adipiscing elit",
		publishDate: rightNow + 20,
	}
	p4 := &Post{
		body:        "sed do eiusmod tempor incididunt",
		publishDate: rightNow + 30,
	}

	f.Append(p1)
	f.Append(p2)
	f.Append(p3)
	f.Append(p4)

	f.Inspect()

	newPost := &Post{
		body:        "This is a new post",
		publishDate: rightNow + 15,
	}
	// вставка + 1
	f.Insert(newPost)
	f.Inspect()

	// удаление - 1
	f.Remove(rightNow + 150)
	f.Inspect()
}
