package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

const (
	// TODO fill this in directly or through environment variable
	// Build a DSN e.g. postgres://username:password@url.com:5432/dbName
	DB_DSN = "postgres://postgres:postgres@localhost:5432/postgres"
)

type User struct {
	ID       int
	Email    string
	Password string
}

func call_db(id int) User {
	fmt.Println(id)
	// Create DB pool
	db, err := sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}

	defer db.Close()

	var myUser User

	userSql := "SELECT id, email, password FROM users WHERE id = $1"

	err = db.QueryRow(userSql, id).Scan(&myUser.ID, &myUser.Email, &myUser.Password)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	fmt.Printf("Hi %s, welcome back!\n", myUser.Email)
	return myUser
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		num, _ := strconv.Atoi(strings.Trim(r.URL.Path, "/"))
		fmt.Println(num)
		fmt.Fprintf(w, "Hello, %q", call_db(num))

	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

	// Create an empty user and make the sql query (using $1 for the parameter)

}
