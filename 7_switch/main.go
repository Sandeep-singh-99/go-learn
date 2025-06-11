package main

import (
	"fmt"
	"time"
)

func main() {
	var day = 2

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown day")
	}


	// multiple condition switch 

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a workday.")
	}


	// type switch

	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("I am an int")
		case string:
			fmt.Println("I am a string")
		case bool:
			fmt.Println("I am a bool")
		default:
			fmt.Println("I am something else")
		}
	}

	whoAmI(42)
}