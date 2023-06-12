package pet

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type uploadResponse struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type PetController struct { // Pet контроллер
	storage PetStorager
}

func NewPetController() *PetController { // конструктор нашего контроллера
	return &PetController{storage: NewPetStorage()}
}

// @Summary create new pet
// @Tags pet
// @ID pet-create-handler
// @Description Create Pet
// @Param body body Pet true "Pet data"
// @Produce json
// @Success 200 {object} Pet
// @Failure 405 "Invalid input"
// @Router /pet [post] PetCreate
func (p *PetController) PetCreate(w http.ResponseWriter, r *http.Request) {
	var pet Pet
	err := json.NewDecoder(r.Body).Decode(&pet) // считываем приходящий json из *http.Request в структуру Pet

	if err != nil { // в случае ошибки отправляем ошибку Bad request code 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	pet = p.storage.Create(pet) // создаем запись в нашем storage

	w.Header().Set("Content-Type", "application/json;charset=utf-8") // выставляем заголовки, что отправляем json в utf8
	err = json.NewEncoder(w).Encode(pet)                             // записываем результат Pet json в http.ResponseWriter

	if err != nil { // отправляем 500 ошибку в случае неудачи
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}

// @Summary get pet by ID
// @Tags pet
// @ID pet-get-by-id-handler
// @Description Get Pet by ID
// @Param id path integer true "petID"
// @Produce json
// @Success 200 {object} Pet
// @Failure 400 "Invalid ID supplied"
// @Failure 404 "Pet not found"
// @Router /get/{petID} [get] PetGetByID
func (p *PetController) PetGetByID(w http.ResponseWriter, r *http.Request) {
	var ( // заранее аллоцируем все необходимые переменные во избежание shadowing
		pet      Pet
		err      error
		petIDRaw string
		petID    int
	)

	petIDRaw = chi.URLParam(r, "petID") // получаем petID из chi router

	petID, err = strconv.Atoi(petIDRaw) // конвертируем в int
	if err != nil {                     // в случае ошибки отправляем код 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	pet, err = p.storage.GetByID(petID) // пытаемся получить Pet по id
	if err != nil {                     // в случае ошибки отправляем Not Found код 404
		http.Error(w, err.Error(), http.StatusNotFound)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	err = json.NewEncoder(w).Encode(pet)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}

// @Summary get delete by ID
// @Tags pet
// @ID pet-delete-by-id-handler
// @Description Delete Pet by ID
// @Param id path integer true "petID"
// @Produce json
// @Success 200
// @Failure 400 "Invalid ID supplied"
// @Failure 404 "Pet not found"
// @Router /get/{petID} [delete] PetDeleteByID
func (p *PetController) PetDeleteByID(w http.ResponseWriter, r *http.Request) {

	petIDRaw := chi.URLParam(r, "petID") // получаем petID из chi router
	petID, err := strconv.Atoi(petIDRaw) // конвертируем в int
	if err != nil {                      // в случае ошибки отправляем код 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
	err = p.storage.Delete(petID)

	if err != nil { // в случае ошибки отправляем код 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}

// @Summary uploads an image
// @Tags pet
// @ID pet-upload-handler
// @Description Create Pet
// @Param id path integer true "petID"
// @Param file formData file true "Pet data"
// @Produce json
// @Success 200 {object} uploadResponse
// @Router /pet/{petID}/uploadImage [post] PetUploadImage
func (p *PetController) PetUploadImage(w http.ResponseWriter, r *http.Request) {

	petIDRaw := chi.URLParam(r, "petID") // получаем petID из chi router
	petID, err := strconv.Atoi(petIDRaw) // конвертируем в int
	if err != nil {                      // в случае ошибки отправляем код 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	r.ParseMultipartForm(10 << 30)

	file, handler, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = p.storage.AddImage(petID, handler.Filename)

	if err != nil { // в случае ошибки отправляем код 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	var response uploadResponse = uploadResponse{Code: http.StatusOK, Type: handler.Filename, Message: "success upload"}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	err = json.NewEncoder(w).Encode(response)

	if err != nil { // в случае ошибки отправляем код 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}

// @Summary Update an existing pet
// @Tags pet
// @ID pet-put-update-handler
// @Param body body Pet true "Pet object that needs to be added to the store"
// @Produce json
// @Failure 400 "Invalid ID supplied"
// @Failure 404 "Pet not found"
// @Failure 405 "Validation exception"
// @Router /pet [put] PetFullUpdate
func (p *PetController) PetFullUpdate(w http.ResponseWriter, r *http.Request) {
	var pet Pet
	err := json.NewDecoder(r.Body).Decode(&pet) // считываем приходящий json из *http.Request в структуру Pet

	if err != nil { // в случае ошибки отправляем ошибку Bad request code 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	pet, err = p.storage.FullUpdate(pet) // создаем запись в нашем storage

	if err != nil { // отправляем 500 ошибку в случае неудачи
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8") // выставляем заголовки, что отправляем json в utf8
	err = json.NewEncoder(w).Encode(pet)                             // записываем результат Pet json в http.ResponseWriter

}

// @Summary Find Pets by status
// @Tags pet
// @ID pet-find-by-status-handler
// @Description Multiple status values can be provided with comma separated strings
// @Param status  query []string true "Status values that need to be considered for filter Available values : available, pending, sold"
// @Produce json
// @Success 200 {object} []Pet
// @Failure 400 "Invalid status value"
// @Router /pet/findByStatus [get] PetFindByStatus
func (p *PetController) PetFindByStatus(w http.ResponseWriter, r *http.Request) {
	var ( // заранее аллоцируем все необходимые переменные во избежание shadowing
		pets []Pet
		err  error
	)

	query := r.URL.Query()
	filters, present := query["status"] //filters=["available","pending","sold"]
	if !present || len(filters) == 0 {
		http.Error(w, "filters not present", http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	pets = p.storage.FindByStatus(filters) // пытаемся получить Pet по id

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	err = json.NewEncoder(w).Encode(pets)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // не забываем прекратить обработку нашего handler (ручки)
	}
}

// @Summary Updates a petin the store with form data
// @Tags pet
// @ID pet-update-handler
// @Param id path integer true "ID of pet that needs to be updated"
// @Param name formData string false "Updated name of the pet"
// @Param status formData string false "Updated status of the pet"
// @Failure 405 "Invalid input"
// @Router /pet/{petID} [post] PetUpdate
func (p *PetController) PetUpdate(w http.ResponseWriter, r *http.Request) {

	petIDRaw := chi.URLParam(r, "petID") // получаем petID из chi router
	petID, err := strconv.Atoi(petIDRaw) // конвертируем в int
	if err != nil {                      // в случае ошибки отправляем код 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	name := r.FormValue("name")
	status := r.FormValue("status")

	ans := uploadResponse{Code: http.StatusOK, Type: "unknown", Message: "Invalid input"}
	data, _ := json.Marshal(ans)
	if name == "" || status == "" {
		http.Error(w, string(data), http.StatusMethodNotAllowed)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

	err = p.storage.Update(name, status, petID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // не забываем прекратить обработку нашего handler (ручки)
	}

}
