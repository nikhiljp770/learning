package main

import "fmt"

func main() {
	n := 5
	isleftup(n)
	isleftdown(n)
	isrightup(n)
	isrightdown(n)
	diamondup(n)
	diamonddown(n)

}
func isleftup(n int) {
	for lines := 1; lines <= 5; lines++ {
		for star := lines; star <= n; star++ {
			fmt.Print("*")
		}
		for hash := 1; hash <= lines; hash++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
func isleftdown(n int) {
	for lines := 1; lines <= n; lines++ {
		for star := 1; star <= lines; star++ {
			fmt.Print("*")
		}
		for hash := lines; hash <= n; hash++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
func isrightup(n int) {
	for lines := 1; lines <= n; lines++ {
		for hash := 1; hash <= lines; hash++ {
			fmt.Print("#")
		}
		for star := lines; star <= n; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func isrightdown(n int) {
	for lines := 1; lines <= 5; lines++ {
		for hash := lines; hash <= n; hash++ {
			fmt.Print("#")
		}
		for star := 1; star <= lines; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func diamondup(n int) {
	for lines := 1; lines <= 5; lines++ {
		for star := lines; star <= n; star++ {
			fmt.Print("*")
		}
		for hash := 1; hash <= lines; hash++ {
			fmt.Print("+")
		}
		for hash := 1; hash <= lines; hash++ {
			fmt.Print("+")
		}
		for star := lines; star <= n; star++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func diamonddown(n int) {
	for lines := 1; lines <= 5; lines++ {
		for hash := 1; hash <= lines; hash++ {
			fmt.Print("#")
		}
		for star := lines; star <= n; star++ {
			fmt.Print("*")
		}
		for hash := lines; hash <= n; hash++ {
			fmt.Print("*")
		}
		for star := 1; star <= lines; star++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
