package database

import (
	"database/sql"
	"fmt"
	"log"
)

// type Database struct {
// 	db *sql.DB
// }

// func NewDatabaseInstace() *Database {
// 	return &Database{ openDatabase()}
// }

func DatabaseInstance() *sql.DB {
	databaseSource := "root:emon@123@tcp(localhost:3306)/jwt-using-gin-gonic"

	db, err := sql.Open("mysql", databaseSource)
	if err != nil {
		log.Fatalf("Failed to open database. The error is %v.", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database. The error is %v.", err)
	}
	fmt.Println("Connected to Maria DB")
	return db
}
