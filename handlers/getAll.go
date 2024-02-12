package handlers

import (
	"Estudos/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetAll(res http.ResponseWriter, req *http.Request) {
	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Error ao buscar os registros: %v", err)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(todos)
}
