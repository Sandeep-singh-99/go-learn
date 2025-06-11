package main

import "fmt"

const age = 30 // constant variable

var isStudent = true // variable

func main() {
	var name = "Sandeep"
	fmt.Println("Hello, " + name + "!")

	fmt.Println("You are", age, "years old.")
	fmt.Println("Are you a student?", isStudent)

	const (
		name1     = "Sandeep"
		age1      = 30
		isStudent1 = true
	)

	fmt.Println("Hello, " + name1 + "!")
	fmt.Println("You are", age1, "years old.")
	fmt.Println("Are you a student?", isStudent1)
}