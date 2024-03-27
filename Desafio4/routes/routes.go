package routes

import (
	"apiClientes/controllers"
	middlawares "apiClientes/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {

	r := mux.NewRouter()

	//Descomentar caso utilizar a aplicação com dados armazenados
	// models.Pessoa = append(models.Pessoa, models.Pessoas{ID: "1", Nome: "John", Sobrenome: "Doe", Endereco: &models.Enderecos{Cidade: "City X", Estado: "State X"}})
	// models.Pessoa = append(models.Pessoa, models.Pessoas{ID: "2", Nome: "Koko", Sobrenome: "Doe", Endereco: &models.Enderecos{Cidade: "City Z", Estado: "State Y"}})

	r.Use(middlawares.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/contato", controllers.GetPeople).Methods("GET")
	r.HandleFunc("/contato/{id}", controllers.GetPerson).Methods("GET")
	r.HandleFunc("/contato/{id}", controllers.CreatePerson).Methods("POST")
	r.HandleFunc("/contato/{id}", controllers.DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
