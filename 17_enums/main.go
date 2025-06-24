package main

import "fmt"

// enumerated types

type MyType int 

const (
	Received MyType = iota
	Pending
	Confirmed
	Delivered
	Cancelled
)

func changeOrderStatus(status MyType) {
	fmt.Println("Order status changed to:", status)
}

func main() {
	changeOrderStatus(Received)
}