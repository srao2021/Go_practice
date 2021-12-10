package api

import (
	"github.com/gorilla/mux"
)

func APIs() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/employee/{id}", crud.GetEmployee).Methods("GET")
	router.HandleFunc("/employee", crud.CreateEmployee).Methods("POST")
	router.HandleFunc("/employee/{id}", crud.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employee/{id}", crud.DeleteEmployee).Methods("DELETE")

	return router
}
