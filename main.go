package main

import (
	"log"
	"net/http"
	"tmanager-app/handler"
	"tmanager-app/middleware"
	"tmanager-app/service"
	"tmanager-app/storage"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {
	db := storage.NewPostgresDB()
	service := service.NewTaskService(db)
	handler := handler.NewTaskHandler(service)

	r := chi.NewRouter()

	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.RecoverMiddleware)
	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", handler.GetTasks)
		r.Post("/", handler.AddTask)
		r.Delete("/{id}", handler.DeleteTask)
		r.Put("/{id}", handler.UpdateTask)
	})

	r.Handle("/*", http.FileServer(http.Dir(".")))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
