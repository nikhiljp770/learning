package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Reverse struct {
	Revstring string `json:"reversed_string"`
}

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
	R := Reverse{Revstring: output}
	log.Println(R)
	log.Println(output)
	js, err := json.Marshal(R)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/reverse", reverseHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
