package main

import "fmt"

func main() {
	age := 20

	if age < 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}


	if age1 := 20; age1 < 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}
}