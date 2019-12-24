package portin

import "context"

// PaymentProcessor defines behavior to process a payment.
type PaymentProcessor interface {
	// Process takes a given payment transaction and process the payment.
	Process(ctx context.Context, newPayment *NewPayment) (*PaymentResult, error)
}
