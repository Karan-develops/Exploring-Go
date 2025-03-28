package main

import "fmt"

type paymenter interface {
	pay(float32) bool
	refund(float32, string)
}

type payment struct {
	gateway paymenter
}

type razorpay struct{}

func (r razorpay) pay(amount float32) bool {
	fmt.Println("Making payment of", amount, "using razorpay")
	return true
}

type stripe struct{}

func (s stripe) pay(amount float32) bool {
	fmt.Println("Making payment of", amount, "using stripe")
	return true
}

type paypal struct{}

func (p paypal) pay(amount float32) bool {
	fmt.Println("Making payment of", amount, "using paypal")
	return true
}

func (p paypal) refund(amount float32, name string) {
	fmt.Println("Refunding through paypal of amount", amount, "in user:", name, "account")
}

func main() {
	// Flexibility to use any gateway without changing the struct code
	// razorpayGw := razorpay{}
	// stripeGw := stripe{}
	paypalGw := paypal{}
	p := payment{
		gateway: paypalGw,
	}
	p.gateway.pay(100)
	p.gateway.refund(100, "Karan")
	// Making payment of 100 using paypal
	// Refunding through paypal of amount 100 in user: Karan account
}

// Without interfaces
/*
type payment struct{}

func (p payment) makePayment(amount float32) {
	razorpayGateway := razorpay{}
	razorpayGateway.pay(amount)

	stripeGateway := stripe{}
	stripeGateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment of", amount, "using razorpay")
}

type stripe struct{}

func (s stripe) pay(amount float32){
	fmt.Println("Making payment of",amount, "using stripe")
}

func main() {
	obj := payment{}
	obj.makePayment(100)
}
*/
