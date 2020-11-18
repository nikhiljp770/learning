package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type CreateMobile struct {
	StoreID      string `json:"store_id"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	CostPrice    string `json:"costprice"`
	SellingPrice string `json:"Sellingprice"`
}

func stores(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	var mobile CreateMobile
	json.Unmarshal(reqBody, &mobile)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("%v", mobile)
	senddata(mobile)

}
func senddata(mobile CreateMobile) {

	db, err := sql.Open("mysql", "root:nikhil@tcp(127.0.0.1:7070)/mobilestore")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	result, err := db.Query("INSERT INTO mobile(store_id,brand,model,cost_price,selling_price) VALUES(?,?,?)", mobile.StoreID, mobile.Brand, mobile.Brand, mobile.CostPrice, mobile.SellingPrice)
	if err != nil {
		panic(err.Error())

	}
	fmt.Println(result)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/mobile", stores).Methods("POST")
	log.Fatal(http.ListenAndServe(":9090", router))
}
