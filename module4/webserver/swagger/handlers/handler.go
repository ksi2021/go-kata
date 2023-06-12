package handlers

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/go-chi/chi/v5"
//	_ "github.com/ksi2021/go-kata/module4/webserver/swagger/repository"
//	. "github.com/ksi2021/go-kata/module4/webserver/swagger/service"
//	"io/ioutil"
//	"log"
//	"os"
//	"strconv"
//
//	"net/http"
//)
//
//var service = MakeService("users.json")
//
//func MainHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("content-type", "application/json")
//	fmt.Fprintf(w, "Hello, World!")
//}
//
//// маршруты для работы с User
//
//// @Summary get all users
//// @ID users-get-handler
//// @Produce json
//// @Success 200 {object} User
//// @Router /users [get] UsersGetHandler
//func UsersGetHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("content-type", "application/json")
//	users, err := service.ListUsers()
//
//	if err != nil {
//		fmt.Fprintf(w, `{"error":`+`"`+err.Error()+`"`+`}`)
//		return
//	}
//	data, _ := json.Marshal(users)
//	w.Write(data)
//}
//
//func UsersPostHandler(w http.ResponseWriter, r *http.Request) {
//
//	password := r.FormValue("password")
//	username := r.FormValue("username")
//	email := r.FormValue("email")
//
//	fmt.Println(password, username, email)
//	if email == "" || password == "" || username == "" {
//		w.WriteHeader(204)
//		return
//	}
//
//	ok := service.CreateUser(password, username, email)
//	if ok {
//		w.WriteHeader(201)
//		fmt.Fprintf(w, `{"status": "create successful"}`)
//		return
//	}
//
//}
//
//func UserHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("content-type", "application/json")
//
//	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
//	user, err := service.GetUser(id)
//
//	if err != nil {
//		w.WriteHeader(400)
//		fmt.Fprintf(w, `{"error":`+`"`+err.Error()+`"`+`}`)
//		return
//	}
//	data, _ := json.Marshal(user)
//	w.Write(data)
//}
//
//// маршруты для работы с файлами
//
//func UploadHandler(w http.ResponseWriter, r *http.Request) {
//	// upload of 10 MB files.
//	r.ParseMultipartForm(10 << 20)
//
//	file, handler, err := r.FormFile("file")
//	if err != nil {
//		fmt.Fprintf(w, `{"error":`+`"`+err.Error()+`"`+`}`)
//		return
//	}
//	defer file.Close()
//
//	tempFile, err := ioutil.TempFile("public", "*")
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fileBytes, err := ioutil.ReadAll(file)
//	if err != nil {
//		fmt.Fprintf(w, `{"error":`+`"`+err.Error()+`"`+`}`)
//		return
//	}
//	// write this byte array to our temporary file
//	tempFile.Write(fileBytes)
//	// return that we have successfully uploaded our file!
//	fmt.Fprintf(w, "Successfully Uploaded File\n")
//
//	tempFile.Close()
//	err = os.Rename(tempFile.Name(), "public/"+handler.Filename)
//	if err != nil {
//		fmt.Fprintf(w, `{"error":`+`"`+err.Error()+`"`+`}`)
//		return
//	}
//
//}
//
//func FilesHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("content-type", "application/json")
//
//	files, err := ioutil.ReadDir("public")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	names := make([]string, len(files))
//	for k, file := range files {
//		names[k] = file.Name()
//	}
//
//	data, err := json.Marshal(names)
//	if err != nil {
//		w.WriteHeader(400)
//
//		fmt.Fprintf(w, `{"error":`+`"`+err.Error()+`"`+`}`)
//	}
//
//	ans := `{"files" :` + string(data) + `}`
//
//	fmt.Fprintf(w, ans)
//}
