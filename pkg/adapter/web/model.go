package web

import (
	"github.com/fernandoocampo/paymentgw/pkg/portin"
)

// paymentProcessorResponse contains data for payment standard response.
type paymentProcessorResponse struct {
	Data *portin.PaymentResult `json:"data,omitempty"`
	Errs []string              `json:"errors,omitempty"`
}

// newErronousPaymentProcessorResponse creates a payment processor response with an error
func newErronousPaymentProcessorResponse(errmsg string) paymentProcessorResponse {
	return paymentProcessorResponse{
		Data: nil,
		Errs: []string{errmsg},
	}
}

// newMultipleErronousPaymentProcessorResponse
func newMultipleErronousPaymentProcessorResponse(paymentResult *portin.PaymentResult) paymentProcessorResponse {
	errorsmessages := make([]string, len(paymentResult.ValidationErrors), len(paymentResult.ValidationErrors))
	for index, err := range paymentResult.ValidationErrors {
		errorsmessages[index] = err.Error()
	}
	return paymentProcessorResponse{
		Data: paymentResult,
		Errs: errorsmessages,
	}
}

// newSuccessPaymentProcessorResponse creates a payment processor response with an error
func newSuccessPaymentProcessorResponse(paymentResult *portin.PaymentResult) paymentProcessorResponse {
	return paymentProcessorResponse{
		Data: paymentResult,
	}
}
