package main

import (
	"fmt"
	"time"
)

// order struct

type Order struct {
	ID        string
	Items     []string
	Total     float64
	status    string
	createdAt time.Time // nano seconds precision
}

func newOrder(id string, items []string, total float64) *Order {

	// initial setup goes here...
	myOrder1 := Order{
		ID:        id,
		Items:     items,
		Total:     total,
	}

	return &myOrder1
}

func (O *Order) changeStatus(status string) {
	O.status = status
}

func (O Order) getTotal() float64 {
	return O.Total
}

func main() {
	order := Order{
		ID:        "12345",
		Items:     []string{"item1", "item2"},
		Total:     29.99,
		status:    "pending",
		createdAt: time.Now(),
	}

	order.changeStatus("shipped")
	fmt.Println(order.getTotal())

	// You can access the fields of the order struct
	fmt.Println("Order ID:", order.ID)
	fmt.Println("Items:", order.Items)
	// fmt.Println("Total:", order.Total)
	fmt.Println("Status:", order.status)
	fmt.Println("Created At:", order.createdAt.Format(time.RFC3339Nano))


	myOrder := Order{
		ID:        "67890",
		Items:     []string{"item3", "item4"},
		Total:     49.99,
		status:    "completed",
		createdAt: time.Now(),
	}


	fmt.Println("My Order ID:", myOrder.ID)
	fmt.Println("My Items:", myOrder.Items)
	fmt.Println("My Total:", myOrder.Total)
	fmt.Println("My Status:", myOrder.status)
	fmt.Println("My Created At:", myOrder.createdAt.Format(time.RFC3339Nano))


	order1 := newOrder("54321", []string{"item5", "item6"}, 19.99)

	fmt.Println("Order1 ID:", order1)


	language := struct{
		Name    string
		Version string
	} {
		Name:    "Go",
		Version: "1.20",
	}

	fmt.Println("Language Name:", language)


}
