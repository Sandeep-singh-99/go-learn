package main

import "fmt"

func changeNum(num *int) {
	*num = 5
	fmt.Println("Inside changeNum:", *num)
}

func main() {
	num := 1

	changeNum(&num)

	fmt.Println("In main:", num)
}