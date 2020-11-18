package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i int = 10
	var f float64 = 6.44
	var str1 string = "abc"
	var str2 string = "10.123"

	fmt.Println(float64(i))
	fmt.Println(int(f))

	newInt, err := strconv.ParseInt(str1, 0, 64)
	if err != nil {
		return
	}
	fmt.Println(newInt)

	newfloat, _ := strconv.ParseFloat(str2, 64)
	fmt.Println(newfloat)
}
