package main

import "fmt"

func main() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// Uncomment the following lines to see an infinite loop
	// for {
	// 	fmt.Println("Infinite loop")

	// }

	for i := 0; i < 3; i++ {
		fmt.Println("Iteration:", i)
	}
}