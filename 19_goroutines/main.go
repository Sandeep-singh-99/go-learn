package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Task", id, "is running")
}

func main() {
	for i := 1; i <= 5; i++ {
		go task(i)
	}

	time.Sleep(time.Second * 2) // Wait for goroutines to finish
	fmt.Println("All tasks are complete")
}