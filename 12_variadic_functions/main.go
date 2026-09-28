package main

import "fmt"


func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func sub(nums ...int) int {
	total := 0

	for _, n := range nums {
		total -= n
	}
	return total
}

func multiply(nums ...int) int {
	mul := 1

	for _, n := range nums {
		mul *= n
	}
	return mul
}


func main() {
	fmt.Println("Variadic Functions Example")
	fmt.Println(1, 2, 3, 4, 5)
	fmt.Println("Hello", "World", "from", "Go")
	fmt.Println("Sum of 1, 2, 3, 4, 5:", sum(1, 2, 3, 4, 5))
	fmt.Println("Sub of 1, 2, 3, 4, 5:", sub(1, 2, 3, 4, 5))
	fmt.Println("Multiply of 1, 2, 3, 4, 5:", multiply(1, 2, 3, 4, 5))
}