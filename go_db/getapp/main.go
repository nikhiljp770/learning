package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type MobileStore struct {
	StoreName   string `json:"store_name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

func getstoreshandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:nikhil@tcp(127.0.0.1:7070)/mobilestore")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	rows, err := db.Query("SELECT store_name, address, phone_number  from stores")
	if err != nil {
		log.Fatal(err)
	}
	var stores []MobileStore
	defer rows.Close()
	for rows.Next() {
		var store MobileStore
		err := rows.Scan(&store.StoreName, &store.Address, &store.PhoneNumber)
		if err != nil {
			log.Fatal(err)
		}
		stores = append(stores, store)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(stores)

	js, err := json.Marshal(stores)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/getstores", getstoreshandler).Methods("GET")
	//router.HandleFunc("/", stores)
	log.Fatal(http.ListenAndServe(":9090", router))
}
