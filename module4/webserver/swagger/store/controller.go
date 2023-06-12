package store

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type OrderController struct { // Order контроллер
	storage OrderStorager
}

func NewOrderController() *OrderController { // конструктор нашего контроллера
	return &OrderController{storage: NewOrderStorage()}
}

// @Summary Place an order for a pet
// @Tags store
// @ID create-order-handler
// @Param body body Order true "order placed for purchasing the pet"
// @Produce json
// @Success 200 {object} Order "successful operation"
// @Failure 400 "Invalid Order"
// @Router /store/order [post] OrderCreate
func (p *OrderController) OrderCreate(w http.ResponseWriter, r *http.Request) {
	var Order Order
	err := json.NewDecoder(r.Body).Decode(&Order) // считываем приходящий json из *http.Request в структуру Order

	if err != nil { // в случае ошибки отправляем ошибку Bad request code 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	Order = p.storage.Create(Order) // создаем запись в нашем storage

	w.Header().Set("Content-Type", "application/json;charset=utf-8") // выставляем заголовки, что отправляем json в utf8
	err = json.NewEncoder(w).Encode(Order)                           // записываем результат Order json в http.ResponseWriter

	if err != nil { // отправляем 500 ошибку в случае неудачи
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}

// @Summary Place an order for a pet
// @Tags store
// @ID get-order-handler
// @Description For valid response try integer IDs with value >= 1 and <= 10. Other values will generated exceptions
// @Param orderID path int true "ID of pet that needs to be fetched"
// @Produce json
// @Success 200 {object} Order "successful operation"
// @Failure 400 "Invalid Order"
// @Failure 404 "Order not found"
// @Router /store/order/{orderID} [get] OrderGetByID
func (p *OrderController) OrderGetByID(w http.ResponseWriter, r *http.Request) {
	var ( // заранее аллоцируем все необходимые переменные во избежание shadowing
		Order   Order
		err     error
		OrderID int
	)

	id := chi.URLParam(r, "orderID") // получаем OrderID из chi router
	OrderID, err = strconv.Atoi(id)

	if err != nil { // в случае ошибки отправляем код 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	Order, err = p.storage.Get(OrderID) // пытаемся получить Order по id
	if err != nil {                     // в случае ошибки отправляем Not Found код 404
		http.Error(w, err.Error(), http.StatusNotFound)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	err = json.NewEncoder(w).Encode(Order)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}

// @Summary Place an order for a pet
// @Tags store
// @Description For valid response try integer IDs with positive integer value. Negative or non-integer values will generate API errors
// @ID delete-order-handler
// @Param orderID path int true "ID of pet that needs to be deleted"
// @Failure 400 "Invalid ID supplied"
// @Failure 404 "Order not found"
// @Router /store/order/{orderID} [delete] OrderDeleteByID
func (p *OrderController) OrderDeleteByID(w http.ResponseWriter, r *http.Request) {
	var ( // заранее аллоцируем все необходимые переменные во избежание shadowing
		err     error
		OrderID int
	)

	id := chi.URLParam(r, "orderID") // получаем OrderID из chi router
	OrderID, err = strconv.Atoi(id)

	if err != nil { // в случае ошибки отправляем код 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	err = p.storage.Delete(OrderID)

	if err != nil { // в случае ошибки отправляем код 400
		http.Error(w, err.Error(), http.StatusNotFound)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}

// @Summary Returns pet inventories by status
// @Tags store
// @Description Returns a map of status codes to quantities
// @ID get-inventory-handler
// @Produce json
// @Success 200 "successful operation"
// @Router /store/order/inventory [get] OrderGetInventory
func (p *OrderController) OrderGetInventory(w http.ResponseWriter, r *http.Request) {

	inventory := make(map[string]int, 0)
	list := p.storage.GetList()
	for _, v := range list {
		inventory[v.Status]++
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	err := json.NewEncoder(w).Encode(inventory)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}
