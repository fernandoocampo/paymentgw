package web

import (
	"context"

	"github.com/fernandoocampo/paymentgw/pkg/portin"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints contains the application capabilities.
type Endpoints struct {
	PaymentProcessorEndpoint endpoint.Endpoint
}

// MakePaymentProcessorEndpoint creates a payment processor endpoint.
func MakePaymentProcessorEndpoint(srv portin.PaymentProcessor) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(portin.NewPayment)
		v, err := srv.Process(ctx, &req)
		if err != nil {
			return newErronousPaymentProcessorResponse(err.Error()), nil
		}
		if v != nil && len(v.ValidationErrors) > 0 {
			return newMultipleErronousPaymentProcessorResponse(v), nil
		}
		return newSuccessPaymentProcessorResponse(v), nil
	}
}

// Process takes a given payment transaction and process the payment.
func (e *Endpoints) Process(ctx context.Context, newPayment *portin.NewPayment) (*portin.PaymentResult, error) {
	resp, err := e.PaymentProcessorEndpoint(ctx, newPayment)
	if err != nil {
		return nil, err
	}
	paymentResp := resp.(paymentProcessorResponse)
	return paymentResp.Data, nil
}
