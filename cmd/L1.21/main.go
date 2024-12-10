package main

import "fmt"

// Интерфейс, который ожидается клиентом
type PaymentProcessor interface {
	ProcessPayment(amount float64)
}

// Существующая система платёжного процессора A
type PayPal struct{}

func (p *PayPal) MakePayment(amount float64) {
	fmt.Printf("PayPal: Processing payment of %.2f\n", amount)
}

// Существующая система платёжного процессора B
type Stripe struct{}

func (s *Stripe) Charge(amount float64) {
	fmt.Printf("Stripe: Charging payment of %.2f\n", amount)
}

// Адаптер для PayPal
type PayPalAdapter struct {
	PayPal *PayPal
}

func (a *PayPalAdapter) ProcessPayment(amount float64) {
	a.PayPal.MakePayment(amount)
}

// Адаптер для Stripe
type StripeAdapter struct {
	Stripe *Stripe
}

func (a *StripeAdapter) ProcessPayment(amount float64) {
	a.Stripe.Charge(amount)
}

func main() {
	// Создаём адаптеры для каждой платёжной системы
	payPalProcessor := &PayPalAdapter{PayPal: &PayPal{}}
	stripeProcessor := &StripeAdapter{Stripe: &Stripe{}}

	// Клиент использует общий интерфейс
	clientsPaymentProcessor := []PaymentProcessor{payPalProcessor, stripeProcessor}

	// Осуществляем платёж через оба процессора
	for _, processor := range clientsPaymentProcessor {
		processor.ProcessPayment(100.0)
	}
}
