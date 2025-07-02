package main

import (
	"math/rand"
	"fmt"
	"time"
)

func processNum(numChan chan int) {

	for nums := range numChan {
		fmt.Println("Processing number:", nums)
		time.Sleep(1 * time.Second) // Simulate some processing time
	}
	// fmt.Println("Processing numbers...", <-numChan)
}


func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

func task(done chan bool) {
	defer func () {
		done <- true
	}()

	fmt.Println("Task is running...")

}


func emailSender(emailChan chan string, done chan bool) {
	defer func() {
		done <- true
	}()
	for email := range emailChan {
		fmt.Println("Sending email to:", email)
	}
}

func main() {
	numChan := make(chan int)

	result := make(chan int)

	emailChan := make(chan string, 100)

	done1 := make(chan bool)
	go emailSender(emailChan, done1)

	emailChan <- "test@example.com"
	emailChan <- "hello@example.com"
	emailChan <- "foo@example.com"

	fmt.Println("Email Channel Length:", len(emailChan))

	go sum (result, 10, 20)

	res := <-result

	fmt.Println("Sum is:", res)


	done := make(chan bool)

	go task(done)

	<-done



	go processNum(numChan)	

	numChan <- 42

	for {
		numChan <- rand.Intn(100)
	}


	// time.Sleep(1 * time.Second)

	// messageChan := make(chan string)

	// messageChan <- "Hello, World!"

	// msg := <-messageChan

	// fmt.Println(msg)
}