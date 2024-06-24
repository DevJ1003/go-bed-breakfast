package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// connect to database
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/bookings")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Success")

	// test my connection
	err = db.Ping()
	if err != nil {
		log.Println("Cannot pinged database")
	}
	log.Println("Pinged database")

	// get rows from table

	// insert a row

	// get rows from table again

	// update a row

	// get rows from table again

	// get one row by id

	// delete a row

	// // get rows from table again
}
