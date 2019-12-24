package application_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fernandoocampo/paymentgw/pkg/domain"
	"github.com/fernandoocampo/paymentgw/pkg/portin"

	"github.com/fernandoocampo/paymentgw/pkg/application"
)

func TestProcess(t *testing.T) {
	ctx := context.TODO()
	t.Run("all_good", func(t *testing.T) {
		service := application.NewBasicPaymentProcessor()
		given := &portin.NewPayment{
			Amount:     5000.0,
			CompanyKey: "12",
			TxID:       "1",
		}
		want := anExpectedValidResult("1", "12", 5000.0)
		got, err := service.Process(ctx, given)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
			t.FailNow()
		}

		want.ID = got.ID
		want.Created = got.Created
		assert.Equal(t, want, got)
	})

	t.Run("all_validations_bad", func(t *testing.T) {
		service := application.NewBasicPaymentProcessor()
		given := &portin.NewPayment{
			Amount:     -5000.0,
			CompanyKey: "",
			TxID:       "",
		}
		want := anExpectedInvalidResult("", "", -5000.0)
		got, err := service.Process(ctx, given)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
			t.FailNow()
		}
		assert.Equal(t, want, got)
	})

	t.Run("all_good_but_internal_error", func(t *testing.T) {
		wantedError := errors.New("fail")
		failFunction := func() error { return errors.New("fail") }
		service := application.NewBasicPaymentProcessorWithFuzzyLogic(failFunction)
		given := &portin.NewPayment{
			Amount:     5000.0,
			CompanyKey: "12",
			TxID:       "1",
		}
		want := given.NewInvalidPaymentResult(portin.NewValidationError(wantedError))
		got, err := service.Process(ctx, given)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
			t.FailNow()
		}

		want.ID = got.ID
		want.Created = got.Created
		assert.Equal(t, want, got)
	})
}

func anExpectedValidResult(txID, companyKey string, amount float32) *portin.PaymentResult {
	payment := &domain.Payment{
		TxID:       txID,
		CompanyKey: companyKey,
		Amount:     amount,
	}
	return portin.NewSuccessPaymentResult(payment)
}

func anExpectedInvalidResult(txID, companyKey string, amount float32) *portin.PaymentResult {
	newpayment := &portin.NewPayment{
		TxID:       txID,
		CompanyKey: companyKey,
		Amount:     amount,
	}
	newerrors := []error{
		errors.New("payment amount cannot be less than zero"),
		errors.New("company key is mandatory"),
		errors.New("a client transaction ID must be provided"),
	}
	return newpayment.NewInvalidPaymentResult(portin.NewValidationErrors(newerrors))
}
