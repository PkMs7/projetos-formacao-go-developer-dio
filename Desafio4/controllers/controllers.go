package controllers

import (
	"apiClientes/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Home Page Modularizado")

}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Pessoa)
}

// GetPerson mostra apenas um contato
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range models.Pessoa {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Pessoas{})
}

// CreatePerson cria um novo contato
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var p models.Pessoas
	_ = json.NewDecoder(r.Body).Decode(&p)
	p.ID = params["id"]
	models.Pessoa = append(models.Pessoa, p)
	json.NewEncoder(w).Encode(models.Pessoa)
}

// DeletePerson deleta um contato
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range models.Pessoa {
		if item.ID == params["id"] {
			models.Pessoa = append(models.Pessoa[:index], models.Pessoa[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(models.Pessoa)
	}
}
