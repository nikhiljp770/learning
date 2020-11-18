package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/palindrome", palindromeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func palindromeHandler(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["string"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}
	key := keys[0]
	result := ispalindrome(key)

	fmt.Fprintf(w, "%v", result)

}
func reverse(key string) string {
	var res string
	var output string
	for i := len(key) - 1; i >= 0; i-- {
		res = string(key[i])
		output = output + res
	}
	return output
}
func ispalindrome(key string) bool {
	res := reverse(key)
	if key == res {
		return true
	}
	return false
}
