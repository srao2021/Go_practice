package db_conn

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest_api/models"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	var employee models.Employee

	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	insertID := InsertEmployee(employee.name, employee.age)

	res := response{
		ID:      insertID,
		Message: "Employee created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

func GetUEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	employee, err := GetEmployee(int64(id))

	if err != nil {
		log.Fatalf("Unable to get employee. %v", err)
	}

	json.NewEncoder(w).Encode(employee)
}

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

	updatedRows := UpdateEmployee(int64(id), employee.name, employee.age)

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

	deletedRows := DeleteEmployee(int64(id))

	msg := fmt.Sprintf("Employee updated successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
