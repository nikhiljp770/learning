package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, value := range nums { // "_ " is to ignore the index
		sum += value
	}
	fmt.Println("sum:", sum)

	for _, num := range nums {
		if num == 3 {
			fmt.Println("index:", num)
		}
	}
	kvs := map[string]string{"1": "mango", "2": "apple", "3": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}
	for i, c := range "Hi" {
		fmt.Println(i, c)
	}

	var a int
	var b int
	for a = 1; a <= 3; a++ {
		for b = 1; b <= 3; b++ {
			if a == 2 && b == 2 {
				break
			}
			fmt.Print(a, " ", b, "\n")
		}
	}
	var i int = 10
	var f float64 = 6.44
	var str1 string = "101"
	var str2 string = "10.123"

	fmt.Println(float64(i))
	fmt.Println(int(f))

	newInt, _ := strconv.ParseInt(str1, 0, 64)
	fmt.Println(newInt)

	newfloat, _ := strconv.ParseFloat(str2, 64)
	fmt.Println(newfloat)
}
