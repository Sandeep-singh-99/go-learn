package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	fmt.Printf("Hello, Sandeep %s\n", "Singh")

	fmt.Println("SANDEEP SINGH")

	fmt.Print("SANDEEP SINGH\n")

	fmt.Println(2 + 5)

	fmt.Printf("2 + 5 = %d\n", 2+5)

	fmt.Printf("Hello, %s", "Sandeep")
	fmt.Println("Welcome")

	// // DON'T use the wrong format verb (e.g., %s for integers):
	fmt.Printf("Age: %s\n", 25) // Compile error or wrong output

	// DO add \n at the end of fmt.Printf:
	fmt.Printf("Hello, %s\n", "Sandeep")
	// DO use the correct format specifiers:
	// %s -> String
	// %d -> Integer (whole number)
	// %.2f -> Float (decimal number, 2 decimal places)
	// %v -> Any value (default format)
	fmt.Printf("Name: %s, Age: %d\n", "Sandeep", 25)

	 //  Best for formatted templates:
    fmt.Printf("Hello, %s! You scored %d points.\n", "Sandeep", 100)
    //  Best for progress / same-line printing:
    fmt.Print("Loading... ")
    fmt.Println("Done!") // Output: Loading... Done!
}
