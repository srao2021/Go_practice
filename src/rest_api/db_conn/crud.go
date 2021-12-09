package db_conn

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	DB_DSN = "postgres://postgres:postgres@localhost:5432/postgres"
)

func InsertEmployee(name string, age int) int {

	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO employee (name, age) VALUES ($1, $2) RETURNING id`

	var id int

	err := db.QueryRow(sqlStatement, name, age).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	return id
}

func createConnection() *sql.DB {

	db, err := sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

func UpdateEmployee(id int, name string, age int) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `UPDATE employee SET name=$2,age=$3 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, name, age)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func DeleteEmployee(id int) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM employee WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
