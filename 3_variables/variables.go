package main

import "fmt"

func main() {
	// var name string = "John"

	// infer type
	// var name = "John"
	var name = "John"

	// boolean variable
	var isStudent = true

	fmt.Println("Hello, " + name + "!")
	fmt.Println("Is student: ", isStudent)

	var age int = 20
	fmt.Println("Age:", age)

	// shorthand syntax

	name1 := "Alice"
	fmt.Println("Hello, " + name1 + "!")

	var name2 string

	name2 = "Bob"
	fmt.Println("Hello, " + name2 + "!")

	var age1 int

	fmt.Print(age1)

	// Grouped declaration (factored block)
	var (
		serverPort int    = 8080
		hostName   string = "localhost"
		debugMode  bool   = false
	)

	fmt.Print(serverPort, "\n", hostName, "\n", debugMode)
}

// := only inside functions
