package web

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/fernandoocampo/paymentgw/pkg/portin"
)

func decodePaymentRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req portin.NewPayment
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
