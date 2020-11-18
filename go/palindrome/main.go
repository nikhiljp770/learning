package main

import "fmt"

func main() {
	A := "liril"
	result := ispalindrome(A)
	fmt.Print(result)

}

func reverse(A string) string {
	var res string
	var output string
	for i := len(A) - 1; i >= 0; i-- {
		res = string(A[i])
		output = output + res
	}
	return output
}
func ispalindrome(A string) bool {
	res := reverse(A)
	if A == res {
		return true
	}
	return false
}
