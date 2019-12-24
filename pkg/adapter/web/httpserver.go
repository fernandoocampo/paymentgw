package web

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

// NewHTTPServer creates a new HTTP Server.
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/payments", httptransport.NewServer(
		endpoints.PaymentProcessorEndpoint,
		decodePaymentRequest,
		encodeResponse,
	))
	return m
}
