package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ksi2021/go-kata/module4/webserver/http/homework/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	port := ":3000"

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handlers.MainHandler)
	r.Get("/users", handlers.UsersGetHandler)
	r.Post("/users", handlers.UsersPostHandler)
	r.Get("/user/{id:[0-9]+}", handlers.UserHandler)
	r.Post("/upload", handlers.UploadHandler)
	r.Get("/public/filename.txt", handlers.FilesHandler)

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
