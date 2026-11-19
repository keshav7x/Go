package main

import "fmt"

type paymentr interface {
	pay(amount float32)
}

type pay struct {
	gateway paymentr
}

func (p pay) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type stripe struct{}
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment with razorpay", amount)
}

func (s stripe) pay(amount float32) {
	fmt.Println("Making payment with Stripe", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32){
	fmt.Println("Making payment with Paypal", amount)
} 


func main() {
	stripePaymentGW := paypal{}
	p1 := pay{
		gateway: stripePaymentGW,
	}
	p1.makePayment(45)
}
