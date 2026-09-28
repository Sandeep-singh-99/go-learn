package main

import "fmt"

func main() {
	// int
	fmt.Println(1 + 1)

	//string
	fmt.Println("Hello, " + "World!")

	// boolean
	fmt.Println(true && false)
	fmt.Println(true || false)

	// float
	fmt.Println(3.14)

	// Integers & Arithmetic
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("7 / 3 =", 7/3) // Truncated to int: 2
	fmt.Println("7 % 3 =", 7%3) // Remainder: 1
	// Floats
	fmt.Println("7.0 / 3.0 =", 7.0/3.0)
	// Strings
	fmt.Println("Hello, " + "World!")
	fmt.Println(`Raw string: no \n escape`)
	// Booleans & Logical Operators
	fmt.Println("true && false:", true && false)
	fmt.Println("true || false:", true || false)
	fmt.Println("!true:", !true)
	fmt.Println("10 > 5:", 10 > 5)
	// Runes (Unicode characters)
	fmt.Println("Rune 'A':", 'A') // 65
	// Type Inspection
	fmt.Printf("Type of 42: %T, type of 3.14: %T\n", 42, 3.14)
}