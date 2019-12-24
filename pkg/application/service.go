package application

import (
	"context"
	"errors"
	"math/rand"

	"github.com/fernandoocampo/paymentgw/pkg/portin"
)

// basicPaymentProcessor implements payment processor use case.
type basicPaymentProcessor struct {
	withLogic  bool
	fuzzyLogic func() error
}

// NewBasicPaymentProcessor creates a basic payment processor implementation.
func NewBasicPaymentProcessor() portin.PaymentProcessor {
	return &basicPaymentProcessor{}
}

// NewBasicPaymentProcessorWithFuzzyLogic creates a basic payment processor implementation.
func NewBasicPaymentProcessorWithFuzzyLogic(fuzzyLogic func() error) portin.PaymentProcessor {
	return &basicPaymentProcessor{
		withLogic:  true,
		fuzzyLogic: fuzzyLogic,
	}
}

// Process takes a given payment transaction and process the payment.
func (b *basicPaymentProcessor) Process(ctx context.Context, newPayment *portin.NewPayment) (*portin.PaymentResult, error) {
	if newPayment == nil {
		return nil, errors.New("given payment cannot be empty")
	}
	validationResult := newPayment.ValidatePayment()

	if validationResult != nil {
		return newPayment.NewInvalidPaymentResult(validationResult), nil
	}

	err := b.applyPayment(newPayment)
	if err != nil {
		applicationError := portin.NewValidationError(err)
		return newPayment.NewInvalidPaymentResult(applicationError), nil
	}

	payment := newPayment.AsPayment()

	return portin.NewSuccessPaymentResult(&payment), nil
}

func (b *basicPaymentProcessor) applyPayment(newPayment *portin.NewPayment) error {
	if newPayment == nil {
		return nil
	}
	if b.withLogic {
		if b.fuzzyLogic != nil {
			return b.fuzzyLogic()
		}
	} else {
		// ...so throw the dices
		dice := rand.Intn(99) + 1
		if dice == 13 { // bad luck
			return errors.New("payment cannot be applied because of an internal error")
		}
	}

	return nil
}
