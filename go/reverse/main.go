package main

import "fmt"

func main() {
	n := "nikhil"
	final := reverse(n)
	fmt.Print(final)
}
func reverse(n string) string {
	var output string
	for i := len(n) - 1; i >= 0; i-- {
		res := string(n[i])
		output = output + res
	}
	return output
}
