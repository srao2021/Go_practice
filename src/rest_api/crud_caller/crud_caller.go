package crud_caller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/srao2021/rest_api/crud"
	"github.com/srao2021/rest_api/models"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	var employee models.Employee

	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	insertID := crud.DBInsertEmployee(employee.Name, employee.Age)

	res := response{
		ID:      int64(insertID),
		Message: "Employee created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

// func GetUEmployee(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	id, err := strconv.Atoi(params["id"])

// 	if err != nil {
// 		log.Fatalf("Unable to convert the string into int.  %v", err)
// 	}

// 	employee, err := DBGetEmployee(int64(id))

// 	if err != nil {
// 		log.Fatalf("Unable to get employee. %v", err)
// 	}

// 	json.NewEncoder(w).Encode(employee)
// }

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	var employee models.Employee

	err = json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := crud.DBUpdateEmployee(id, employee.Name, employee.Age)

	msg := fmt.Sprintf("Employee updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := crud.DBDeleteEmployee(id)

	msg := fmt.Sprintf("Employee updated successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
