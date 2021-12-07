package main

import (
	"database/sql"
	"fmt"
	"log"

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

func main() {

	// Create DB pool
	db, err := sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	// Create an empty user and make the sql query (using $1 for the parameter)
	var myUser User

	deleteStmt := `delete from "users" where id=2`
	_, e1 := db.Exec(deleteStmt)
	CheckError(e1)

	insertStmt := `insert into "users"("id", "email", "password") values(2, 'smith@123.com', 'smith')`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	updateStmt := `update "users" set email='smith@1234.com' where id =$1`
	_, e2 := db.Exec(updateStmt, 2)
	CheckError(e2)

	userSql := "SELECT id, email, password FROM users WHERE id = $1"

	err = db.QueryRow(userSql, 2).Scan(&myUser.ID, &myUser.Email, &myUser.Password)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	fmt.Printf("Hi %s, welcome back!\n", myUser.Email)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
