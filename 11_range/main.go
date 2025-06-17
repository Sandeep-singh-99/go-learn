package main

import "fmt"

// iterating over data structures in Go

func main() {
	nums := []int{1, 2, 3, 4, 5}

	 for i, num := range nums {
	// i is the index, num is the value
	println("Index:", i, "Value:", num)
  }

  
	// unicode code point rune
	// starting byte of rune
	// 300 -> 1 byte, 2 byte

   for j, char := range "Hello" {
		fmt.Println(j, string(char))
   }
}