package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:nikhil@tcp(127.0.0.1:7070)/mobile_store")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO stores(store_name,address,phone_number) VALUES('nikhil','banglore',9743289770)")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Print("inserted succesfully")
}
