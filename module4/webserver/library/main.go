package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/ksi2021/go-kata/module4/webserver/library/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title Library Rest API
// @version 1.0
// @description copy of the swagger Library

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
func main() {

	db, err := DBConnect()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Author{}, &Book{}, &User{})

	FillAuthors(db, 10)
	FillBooks(db, 100)
	FillUsers(db, 50)

	port := ":8080"
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	controller := NewController(db)

	r.Get("/getUsers", controller.GetUsers)
	r.Get("/getAuthors", controller.GetAuthors)
	r.Get("/getBooks", controller.GetBooks)

	r.Post("/createBook", controller.CreateBook)
	r.Post("/createAuthor", controller.CreateAuthor)
	r.Post("/takeBook", controller.TakeBook)

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
