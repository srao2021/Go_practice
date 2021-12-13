package api

import (
	"github.com/gorilla/mux"
	"github.com/srao2021/rest_api/crud_caller"
)

func APIs() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/employee", crud_caller.CreateEmployee).Methods("POST")
	router.HandleFunc("/employee/{id}", crud_caller.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employee/{id}", crud_caller.DeleteEmployee).Methods("DELETE")

	return router
}
