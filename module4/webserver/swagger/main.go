package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/ksi2021/go-kata/module4/webserver/swagger/docs"
	. "github.com/ksi2021/go-kata/module4/webserver/swagger/pet"
	. "github.com/ksi2021/go-kata/module4/webserver/swagger/store"
	. "github.com/ksi2021/go-kata/module4/webserver/swagger/user"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title Swagger Rest API
// @version 1.0
// @description copy of the swagger PetStore

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
func main() {
	port := ":8080"

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// pets paths
	petController := NewPetController()
	r.Post("/pet", petController.PetCreate)
	r.Get("/pet/{petID:[0-9]+}", petController.PetGetByID)
	r.Delete("/pet/{petID:[0-9]+}", petController.PetDeleteByID)
	r.Post("/pet/{petID:[0-9]+}/uploadImage", petController.PetUploadImage)
	r.Get("/pet/findByStatus", petController.PetFindByStatus)
	r.Post("/pet/{petID:[0-9]+}", petController.PetUpdate)
	r.Put("/pet", petController.PetFullUpdate)

	// user paths
	userController := NewUserController()
	r.Post("/user", userController.UserCreate)
	r.Post("/user/createWithList", userController.UserCreateWithList)
	r.Post("/user/createWithArray", userController.UserCreateWithList)
	r.Get("/user/{Username:[a-zA-Z]+}", userController.UserGetByUsername)
	r.Delete("/user/{Username:[a-zA-Z]+}", userController.UserDeleteByUsername)
	r.Put("/user/{Username:[a-zA-Z]+}", userController.UserUpdate)

	// store paths
	orderController := NewOrderController()
	r.Get("/store/inventory", orderController.OrderGetInventory)
	r.Get("/store/order/{orderID:[0-9]+}", orderController.OrderGetByID)
	r.Delete("/store/order/{orderID:[0-9]+}", orderController.OrderDeleteByID)
	r.Post("/store/order", orderController.OrderCreate)

	// swagger path
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	// Запуск веб-сервера в отдельном горутине
	go func() {
		log.Println(fmt.Sprintf("server started on port %s ", port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Ожидание сигнала для начала завершения работы
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	// Установка тайм-аута для завершения работы
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
