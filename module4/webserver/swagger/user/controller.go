package user

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type uploadResponse struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type UserController struct { // User контроллер
	storage UserStorager
}

func NewUserController() *UserController { // конструктор нашего контроллера
	return &UserController{storage: NewUserStorage()}
}

// @Summary Create user
// @Tags user
// @ID user-create-handler
// @Description This can only be done by the logged in user.
// @Param body body User true "Created user object"
// @Produce json
// @Failure default {string} successful operation "successful operation"
// @Failure 404  "Pet not found"
// @Router /user [post] UserCreate
func (p *UserController) UserCreate(w http.ResponseWriter, r *http.Request) {
	var User User
	err := json.NewDecoder(r.Body).Decode(&User) // считываем приходящий json из *http.Request в структуру User

	if err != nil { // в случае ошибки отправляем ошибку Bad request code 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	User = p.storage.Create(User) // создаем запись в нашем storage

	w.Header().Set("Content-Type", "application/json;charset=utf-8") // выставляем заголовки, что отправляем json в utf8
	err = json.NewEncoder(w).Encode(User)                            // записываем результат User json в http.ResponseWriter

	if err != nil { // отправляем 500 ошибку в случае неудачи
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}

// @Summary Create with list user
// @Tags user
// @ID user-create-with-list-handler
// @Description This can only be done by the logged in user.
// @Param body body []User true "List of user object"
// @Produce json
// @Failure default {string} successful operation "successful operation"
// @Failure 404 "Pet not found"
// @Router /user/createWithList [post] UserCreateWithList
func (p *UserController) UserCreateWithList(w http.ResponseWriter, r *http.Request) {
	var Users []User
	err := json.NewDecoder(r.Body).Decode(&Users) // считываем приходящий json из *http.Request в структуру User

	if err != nil { // в случае ошибки отправляем ошибку Bad request code 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	err = p.storage.CreateWithList(Users) // создаем запись в нашем storage

	if err != nil { // отправляем 500 ошибку в случае неудачи
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	ans := uploadResponse{Code: 200, Type: "unknown", Message: "OK"}
	w.Header().Set("Content-Type", "application/json;charset=utf-8") // выставляем заголовки, что отправляем json в utf8
	err = json.NewEncoder(w).Encode(ans)                             // записываем результат User json в http.ResponseWriter

}

// @Summary Get user
// @Tags user
// @ID user-get-by-username-handler
// @Description This can only be done by the logged in user.
// @Param Username path string true "The name that needs to be fetched"
// @Produce json
// @Success 200 {object} User "successful operation"
// @Failure 400 "Invalid username supplied"
// @Failure 404 "User not found"
// @Router /user/{Username} [get] UserGetByUsername
func (p *UserController) UserGetByUsername(w http.ResponseWriter, r *http.Request) {
	var ( // заранее аллоцируем все необходимые переменные во избежание shadowing
		User     User
		err      error
		Username string
	)

	Username = chi.URLParam(r, "Username") // получаем UserID из chi router

	if Username == "" { // в случае ошибки отправляем код 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	User, err = p.storage.GetByUsername(Username) // пытаемся получить User по id
	if err != nil {                               // в случае ошибки отправляем Not Found код 404
		http.Error(w, err.Error(), http.StatusNotFound)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	err = json.NewEncoder(w).Encode(User)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}

// @Summary Delete user
// @Tags user
// @ID user-delete-by-username-handler
// @Description This can only be done by the logged in user.
// @Param Username path string true "The name that needs to be deleted"
// @Failure 400 "Invalid username supplied"
// @Failure 404 "User not found"
// @Router /user/{Username} [post] UserDeleteByUsername
func (p *UserController) UserDeleteByUsername(w http.ResponseWriter, r *http.Request) {

	Username := chi.URLParam(r, "Username") // получаем UserID из chi router
	if Username == "" {                     // в случае ошибки отправляем код 400
		http.Error(w, "Invalid username supplied", http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	err := p.storage.Delete(Username)

	if err != nil { // в случае ошибки отправляем код 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}

// @Summary Update user
// @Tags user
// @ID user-update-by-username-handler
// @Description This can only be done by the logged in user.
// @Param Username path string true "The name that needs to be updated"
// @Param body body User true "Updated user object"
// @Failure 400 "Invalid username supplied"
// @Failure 404 "User not found"
// @Router /user/{Username} [put] UserDeleteByUsername
func (p *UserController) UserUpdate(w http.ResponseWriter, r *http.Request) {

	var User User

	Username := chi.URLParam(r, "Username") // получаем UserID из chi router
	if Username == "" {                     // в случае ошибки отправляем код 400
		http.Error(w, "Invalid username supplied", http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	err := json.NewDecoder(r.Body).Decode(&User) // считываем приходящий json из *http.Request в структуру User

	if err != nil { // в случае ошибки отправляем ошибку Bad request code 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	User, err = p.storage.Update(Username, User) // создаем запись в нашем storage

	if err != nil { // отправляем 500 ошибку в случае неудачи
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8") // выставляем заголовки, что отправляем json в utf8
	err = json.NewEncoder(w).Encode(User)                            // записываем результат User json в http.ResponseWriter

}
