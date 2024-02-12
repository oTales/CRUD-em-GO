package main

import (
	"Estudos/configs"
	"Estudos/handlers"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {

	err := configs.Load()

	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()

	router.Get("/todos", handlers.GetAll)
	router.Get("/todos/{id}", handlers.Get)
	router.Post("/todos", handlers.Create)
	router.Put("/todos/{id}", handlers.Update)
	router.Delete("/todos/{id}", handlers.Delete)

	err = http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router)
	if err != nil {
		log.Printf("Error ao iniciar o servidor: %v", err)
		return
	}
}
