package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	output := arraysum(a)
	fmt.Println(output)
	output2 := average(a)
	fmt.Println(output2)
}

func arraysum(a [5]int) int {
	var sum int
	for i := 0; i < len(a); i++ {
		//fmt.Println(a[i])
		sum = sum + a[i]
	}
	return sum
}
func average(a [5]int) int {
	var sum int
	var average int
	var count int
	count = len(a)
	sum = arraysum(a)
	average = sum / count
	//fmt.Print(average)
	return average
}
