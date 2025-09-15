package main

import (
	"fmt"
)

// PaymentStrategy defines a common interface
type PayementStrategy interface {
	Pay(amount float64)
}

type CreditCardPayment struct {
	Name   string
	CardNo string
}

func (c *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("paid %.2f using Credit card [%s]\n", amount, c.CardNo)
}

type PaypalPayment struct {
	Email string
}

func (p *PaypalPayment) Pay(amount float64) {
	fmt.Printf("Paid %.2f using paypal account [%s]]\n", amount, p.Email)
}

type PaymentProcessor struct {
	strategy PayementStrategy
}

func (pp *PaymentProcessor) SetStrategy(strategy PayementStrategy) {
	pp.strategy = strategy
}

func (pp *PaymentProcessor) Checkout(amount float64) {
	pp.strategy.Pay(amount)
}

func main() {
	processor := &PaymentProcessor{}

	processor.SetStrategy(&CreditCardPayment{Name: "Shubham", CardNo: "1234-5678"})
	processor.Checkout(250.75) // output - Paid 250.75 using Credit Card [1234-5678]

	processor.SetStrategy(&PaypalPayment{Email: "shubham@example.com"})
	processor.Checkout(99.99) // Paid 99.99 using PayPal account [shubham@example.com]
}
