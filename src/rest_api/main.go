package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/srao2021/rest_api/api"
)

func main() {

	//id := crud.DBInsertEmployee("abc", 1)
	//fmt.Println(id)
	r := api.APIs()

	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
