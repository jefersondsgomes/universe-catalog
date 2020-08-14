package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../models"
	"../repositories"

	"github.com/gorilla/mux"
)

func GetAstros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	astros := repositories.GetAll()

	if len(astros) == 0 {
		w.WriteHeader(204)
		return
	}

	json.NewEncoder(w).Encode(astros)
}

func GetAstro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(400)
		return
	}

	astro := repositories.Get(id)

	if astro.Id < 1 {
		w.WriteHeader(204)
		return
	}

	json.NewEncoder(w).Encode(astro)
}

func CreateAstro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var astro models.Astro
	_ = json.NewDecoder(r.Body).Decode(&astro)
	result := repositories.Create(astro)

	w.WriteHeader(204)
	json.NewEncoder(w).Encode(result)
}

func UpdateAstro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func DeleteAstro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}