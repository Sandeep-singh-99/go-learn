package main

import "fmt"

// numbered sequence of specific length

func main() {
	var nums [5]int
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4

	fmt.Println(len(nums)) // length of the array

	fmt.Println(nums) // print the array

	var vals [5]string
	vals[0] = "one"

	fmt.Println(vals) // print the array

	// to declare and initialize an array in one line
	nums1 := [3]int{1, 2, 3}

	fmt.Println(nums1) // print the array

	// 2d array
	nums2 := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(nums2) // print the 2d array

	// - fixed size, that is predictable
	// - fast access to elements by index
	// - memory is contiguous, which can be more efficient for CPU cache
	// - not flexible, cannot change size after declaration
	// constant time access

}
