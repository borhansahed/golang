package main

import "fmt"

// Good code

// In golang there is no inheritance and implements
// So when you struct methods parameter and name same as interface method then golang compiler implicitly implements the interface to the struct

type paymenter interface {
	pay (amount int)
}

// Payment struct
type payment struct {
	gateway paymenter
}

func (p payment) makePayment (amount int) {
	p.gateway.pay(amount)
}


// Paypal struct
type paypal struct {}

func (p paypal) pay (amount int) {

	// logic to make payment

	fmt.Println("Making payment paypal", amount)
}
// Stripe struct
type stripe struct {}

func (p stripe) pay (amount int) {

	// logic to make payment

	fmt.Println("Making payment stripe", amount)
}

func main () {

	stripePayment := stripe{}
	newPayment := payment{
		gateway: stripePayment ,
	}

	newPayment.makePayment(100)
}
// Bad Code 

// Payment struct
// type payment struct {}

// func (p payment) makePayment (amount int) {
       
// 	paypal := paypal{}

// 	paypal.pay(amount)
// }


// Paypal struct

// type paypal struct {}

// func (p paypal) pay (amount int) {

// 	// logic to make payment

// 	fmt.Println("Making payment", amount)
// }
// func main () {
// 	newPayment := payment{}

// 	newPayment.makePayment(100)
// }