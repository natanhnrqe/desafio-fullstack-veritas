package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {

	InitDB()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders: []string{"Content-Type"},
}))

	r.Get("/tasks", GetTasks)
	r.Post("/tasks", CreateTask)
	r.Put("/tasks/{id}", UpdateTask)
	r.Delete("/tasks/{id}", DeleteTask)

	
	log.Println("Rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
