package portin_test

import (
	"errors"
	"testing"

	"github.com/fernandoocampo/paymentgw/pkg/portin"
	"github.com/stretchr/testify/assert"
)

func TestValidatePayment(t *testing.T) {
	tests := []struct {
		testcase string
		given    portin.NewPayment
		want     *portin.ValidationError
	}{
		{"valid", newPayment("1", "12", 23500.0), nil},
		{"empty_negative", newPayment("", "", -1), anEmptyNegativeValidationResult()},
		{"empty_maximum_reached", newPayment("", "", 5000002.0), anEmptyMaxValidationResult()},
		{"empty_txid", newPayment("", "12", 1000), anEmptyTxIDValidationResult()},
		{"empty_companey_key", newPayment("1", "", 1000), anEmptyCompanyKeyValidationResult()},
		{"maximum_reached", newPayment("1", "12", 5000002.0), aMaxValidationResult()},
		{"negative", newPayment("1", "12", -1), aNegativeValidationResult()},
	}
	for _, test := range tests {
		got := test.given.ValidatePayment()
		assert.Equal(t, test.want, got)
	}
}

func anEmptyNegativeValidationResult() *portin.ValidationError {
	newerrors := []error{
		errors.New("payment amount cannot be less than zero"),
		errors.New("company key is mandatory"),
		errors.New("a client transaction ID must be provided"),
	}
	return portin.NewValidationErrors(newerrors)
}

func anEmptyMaxValidationResult() *portin.ValidationError {
	newerrors := []error{
		errors.New("maximum payment amount allowed is $5000001"),
		errors.New("company key is mandatory"),
		errors.New("a client transaction ID must be provided"),
	}
	return portin.NewValidationErrors(newerrors)
}

func anEmptyTxIDValidationResult() *portin.ValidationError {
	newerrors := []error{
		errors.New("a client transaction ID must be provided"),
	}
	return portin.NewValidationErrors(newerrors)
}

func anEmptyCompanyKeyValidationResult() *portin.ValidationError {
	newerrors := []error{
		errors.New("company key is mandatory"),
	}
	return portin.NewValidationErrors(newerrors)
}

func aMaxValidationResult() *portin.ValidationError {
	newerrors := []error{
		errors.New("maximum payment amount allowed is $5000001"),
	}
	return portin.NewValidationErrors(newerrors)
}

func aNegativeValidationResult() *portin.ValidationError {
	newerrors := []error{
		errors.New("payment amount cannot be less than zero"),
	}
	return portin.NewValidationErrors(newerrors)
}

func newPayment(txID, companyKey string, amount float32) portin.NewPayment {
	return portin.NewPayment{
		TxID:       txID,
		CompanyKey: companyKey,
		Amount:     amount,
	}
}
