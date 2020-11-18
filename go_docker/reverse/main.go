package main

import (
	"fmt"
	"log"
	"net/http"
)

func reverseHandler(w http.ResponseWriter, r *http.Request) {
	var output string
	keys, ok := r.URL.Query()["string"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}
	key := keys[0]
	for i := len(key) - 1; i >= 0; i-- {
		res := string(key[i])
		output = output + res
	}
	log.Println("Url Param 'key' is: " + string(key))
	fmt.Fprintf(w, output)
}

func main() {
	http.HandleFunc("/reverse", reverseHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
