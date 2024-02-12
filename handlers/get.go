package handlers

import (
	"Estudos/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func Get(res http.ResponseWriter, req *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(req, "id"))

	if err != nil {
		log.Printf("Error ao fazer a conversao do id: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	todo, err := models.Get(int64(id))

	if err != nil {
		log.Printf("Error ao buscar registros %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(todo)
}
