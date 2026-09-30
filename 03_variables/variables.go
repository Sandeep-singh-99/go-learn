package main

import "fmt"

var x = 10

func main() {
	x := 20
	if true {
		x := 30
		fmt.Println(x)
	}
	fmt.Println(x)
}