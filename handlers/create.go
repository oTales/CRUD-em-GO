package handlers

import (
	"Estudos/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(res http.ResponseWriter, req *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(req.Body).Decode(&todo)

	if err != nil {
		log.Printf("Error ao fazer decoding do json", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"error":   true,
			"message": fmt.Sprintf("Erro ao inserir o todo: %v", err),
		}
	}

	if err == nil {
		resp = map[string]any{
			"error":   false,
			"message": fmt.Sprintf("Todo inserido com sucesso! id: %d", id),
		}
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(resp)
}
