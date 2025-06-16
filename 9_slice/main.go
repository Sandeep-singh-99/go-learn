package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in go
// + useful methods

func main() {

	// uninitialized slice is nil

	var nums []int

	fmt.Println(nums)        // []
	fmt.Println(nums == nil) // true
	fmt.Println(len(nums))   // 0

	var num1 = make([]int, 5, 10) // initialized with 5 elements, all zero values
	fmt.Println(num1)             // [0 0 0 0 0]

	// capacity of a slice is the number of elements it can hold before needing to allocate more memory
	fmt.Println(cap(num1)) // 5

	num1 = append(num1, 1, 2, 3, 4) // appending elements
	num1[5] = 5                     // modifying an element
	fmt.Println(num1)               // [0 0 0 0 0 1 2 3 4 5]

	// copy function
	num2 := make([]int, len(num1))
	copy(num2, num1)  // copying elements from num1 to num2
	fmt.Println(num2) // [0 0 0 0 0 1 2 3 4 5]

	// slice operators
	var nums3 = []int{1, 2, 3}
	fmt.Println(nums3[0:1]) // [1 2 3]
	fmt.Println(nums3[:2])  // [2 3]

	fmt.Println(slices.Equal(num2, nums3))

}
