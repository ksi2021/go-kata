package main

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Controller struct {
	Store *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{Store: db}
}

// @Summary GetBooks
// @ID books-get-handler
// @Description Получение всех книг вместе с авторами.
// @Produce json
// @Success 200 {object} Book "successful operation"
// @Failure 500
// @Router /getBooks [get] GetBooks
func (c *Controller) GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book

	c.Store.Find(&books)
	c.Store.Preload("Author").Preload("User").Order("id asc").Find(&books)

	w.Header().Set("Content-Type", "application/json;charset=utf-8") // выставляем заголовки, что отправляем json в utf8
	err := json.NewEncoder(w).Encode(books)                          // записываем результат User json в http.ResponseWriter

	if err != nil { // отправляем 500 ошибку в случае неудачи
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}

// @Summary GetAuthors
// @ID authors-get-handler
// @Description Получение авторов вместе с их книгами.
// @Produce json
// @Success 200 {object} Author "successful operation"
// @Failure 500
// @Router /getAuthors [get] GetAuthors
func (c *Controller) GetAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []Author

	c.Store.Find(&authors)
	c.Store.Preload("Books").Find(&authors)

	w.Header().Set("Content-Type", "application/json;charset=utf-8") // выставляем заголовки, что отправляем json в utf8
	err := json.NewEncoder(w).Encode(authors)                        // записываем результат User json в http.ResponseWriter

	if err != nil { // отправляем 500 ошибку в случае неудачи
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}

// @Summary Create new Author
// @ID author-create-handler
// @Param name formData string true "Author name"
// @Failure 400 "Invalid input"
// @Failure 500 "Invalid data"
// @Success 200 {object} Author "successful operation"
// @Router /createAuthor [post] CreateAuthor
func (c *Controller) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author Author

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "invalid data", 400)
		return
	}

	author.Name = name
	result := c.Store.Create(&author)
	if result.Error != nil {
		http.Error(w, "invalid data", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8") // выставляем заголовки, что отправляем json в utf8
	err := json.NewEncoder(w).Encode(author)                         // записываем результат User json в http.ResponseWriter

	if err != nil { // отправляем 500 ошибку в случае неудачи
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}

// @Summary GetUsers
// @ID users-get-handler
// @Description Получение пользователей вместе с книгами.
// @Produce json
// @Success 200 {object} User "successful operation"
// @Failure 500
// @Router /getUsers [get] GetUsers
func (c *Controller) GetUsers(w http.ResponseWriter, r *http.Request) {

	var users []User
	//var books []Book

	//c.Store.Find(&books)
	//c.Store.Preload("Author").Find(&books)
	c.Store.Find(&users)
	c.Store.Preload("RentedBooks.Author").Find(&users)
	//c.Store.Model(&users).Association("Books").Find(&books)

	w.Header().Set("Content-Type", "application/json;charset=utf-8") // выставляем заголовки, что отправляем json в utf8
	err := json.NewEncoder(w).Encode(users)                          // записываем результат User json в http.ResponseWriter

	if err != nil { // отправляем 500 ошибку в случае неудачи
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}

// @Summary Create new book
// @ID book-create-handler
// @Param title formData string true "book title"
// @Param author_id formData string true "Author ID"
// @Failure 400 "Invalid input"
// @Failure 500 "Invalid data"
// @Success 200 {object} Book "successful operation"
// @Router /createBook [post] CreateBook
func (c *Controller) CreateBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	id := r.FormValue("author_id")
	if title == "" || id == "" {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	author_id := uint(ID)

	var author Author

	result := c.Store.First(&author, author_id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	var book Book
	book.Title = title
	book.AuthorID = author_id

	result = c.Store.Create(&book)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	c.Store.Preload("Author").Last(&book)
	w.Header().Set("Content-Type", "application/json;charset=utf-8") // выставляем заголовки, что отправляем json в utf8
	err = json.NewEncoder(w).Encode(book)                            // записываем результат User json в http.ResponseWriter

	if err != nil { // отправляем 500 ошибку в случае неудачи
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}

// @Summary Take a book
// @ID book-take-handler
// @Param book_id formData string true "Book ID"
// @Param user_id formData string true "User ID"
// @Failure 400 "Invalid input"
// @Failure 500 "Invalid data"
// @Success 200 {object} Book "successful operation"
// @Router /takeBook [post] TakeBook
func (c *Controller) TakeBook(w http.ResponseWriter, r *http.Request) {
	book_id_string := r.FormValue("book_id")
	user_id_string := r.FormValue("user_id")
	if book_id_string == "" || user_id_string == "" {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	book_id, err := strconv.Atoi(book_id_string)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	user_id, err := strconv.Atoi(user_id_string)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
	bookID := uint(book_id)
	userID := uint(user_id)
	var user User
	var book Book

	result := c.Store.First(&user, userID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	result = c.Store.First(&book, bookID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	if book.UserID > 0 {
		http.Error(w, "The book is busy", http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	result = c.Store.Model(&book).Update("user_id", userID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	c.Store.Preload("Author").Last(&book)
	book.User = user
	w.Header().Set("Content-Type", "application/json;charset=utf-8") // выставляем заголовки, что отправляем json в utf8
	err = json.NewEncoder(w).Encode(book)                            // записываем результат User json в http.ResponseWriter

	if err != nil { // отправляем 500 ошибку в случае неудачи
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}
