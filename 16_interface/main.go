package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct{
	gateway paymenter
	gateway2 paymenter

}

func (p payment) makePayment(amount float32) {
	
	p.gateway.pay(amount)
	p.gateway2.pay(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	fmt.Println("Payment made using Razorpay:", amount)
}


type stripe struct {

}

func (s stripe) pay(amount float32) {
	fmt.Println("Payment made using Stripe:", amount)
}

func main() {
	// Create a payment instance with both gateways
	// and make a payment
	razorpayGateway := razorpay{}
	stripeGateway := stripe{}
	p := payment{
	gateway:  stripeGateway,
	gateway2: razorpayGateway,
	}
	p.makePayment(100.50)
}