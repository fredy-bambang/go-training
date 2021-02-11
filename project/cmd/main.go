package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rickywinata/go-training/project/internal/db/inmemory"
	"github.com/rickywinata/go-training/project/internal/handlers"
	"github.com/rickywinata/go-training/project/internal/repositories"
	"github.com/rickywinata/go-training/project/internal/services"
)

func main() {
	repo := &inmemory.InmemUserRepository{
		Data: []*repositories.User{},
	}
	svc := services.NewUserService(repo)

	r := chi.NewRouter()
	r.Post("/users", handlers.CreateUserHandler(svc).ServeHTTP)
	r.Get("/users/{id}", nil)

	log.Println("Listening on :8080 ...")
	http.ListenAndServe(":8080", r)
}
