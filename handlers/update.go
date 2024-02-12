package handlers

import (
	"Estudos/models"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func Update(res http.ResponseWriter, req *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(req, "id"))

	if err != nil {
		log.Printf("Error ao fazer a conversao do id: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo models.Todo

	err = json.NewDecoder(req.Body).Decode(&todo)

	if err != nil {
		log.Printf("Error ao fazer decoding do json: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(int64(id), todo)

	if err != nil {
		log.Printf("Error ao fazer update do todo: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram atualizados mais %d registros", rows)
	}

	resp := map[string]any{
		"error":   false,
		"message": fmt.Sprintf("Todo atualizado com sucesso! id: %d", id),
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(resp)
}
