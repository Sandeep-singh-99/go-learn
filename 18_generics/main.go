package main

import "fmt"

func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type Stack[T any] struct {
	elements []T
}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	strs := []string{"hello", "world"}

	printSlice(nums)
	printSlice(strs)

	myStack := Stack[int]{
		elements: []int{10, 20, 30},
	}

	fmt.Println(myStack)
}