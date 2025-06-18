package main

import "fmt"


func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println("Variadic Functions Example")
	fmt.Println(1, 2, 3, 4, 5)
	fmt.Println("Hello", "World", "from", "Go")
	fmt.Println("Sum of 1, 2, 3, 4, 5:", sum(1, 2, 3, 4, 5))
}