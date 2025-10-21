package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

// ---- Different gateways ----
type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Making payment using Stripe =", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("Making payment using PayPal =", amount)
}

// ---- Main ----
func main() {
	// Stripe payment
	stripeGateway := stripe{}
	p1 := payment{gateway: stripeGateway}
	p1.makePayment(500)

	// PayPal payment
	paypalGateway := paypal{}
	p2 := payment{gateway: paypalGateway}
	p2.makePayment(800)
}
