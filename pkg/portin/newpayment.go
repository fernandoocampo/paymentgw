package portin

import (
	"errors"
	"strings"
	"time"

	"github.com/fernandoocampo/paymentgw/pkg/domain"
)

// NewPayment contains data related to new payments.
type NewPayment struct {
	TxID       string  `json:"transaction_id"`
	CompanyKey string  `json:"company_key"`
	Amount     float32 `json:"amount"`
}

// PaymentResult contains data related to transaction result.
type PaymentResult struct {
	ID               string     `json:"confirmation_id,omitempty"`
	TxID             string     `json:"transaction_id"`
	CompanyKey       string     `json:"company_key"`
	Amount           float32    `json:"amount"`
	Success          bool       `json:"success"`
	ValidationErrors []error    `json:"-"`
	Created          *time.Time `json:"created_at,omitempty"`
}

// ValidationError define validation error result over model data.
type ValidationError struct {
	errors []error
}

// AsPayment converts new payment to payment.
func (n NewPayment) AsPayment() domain.Payment {
	return domain.NewPayment(n.TxID, n.CompanyKey, n.Amount)
}

// NewSuccessPaymentResult creates a successfully payment result
func NewSuccessPaymentResult(payment *domain.Payment) *PaymentResult {
	createdAt := payment.Created
	return &PaymentResult{
		Success:    true,
		TxID:       payment.TxID,
		Amount:     payment.Amount,
		CompanyKey: payment.CompanyKey,
		Created:    &createdAt,
		ID:         payment.ID,
	}
}

// NewInvalidPaymentResult creates an invalid payment result from newpayment
func (n NewPayment) NewInvalidPaymentResult(validationError *ValidationError) *PaymentResult {
	return &PaymentResult{
		Success:          false,
		TxID:             n.TxID,
		Amount:           n.Amount,
		CompanyKey:       n.CompanyKey,
		ValidationErrors: validationError.errors,
	}
}

// NewValidationErrors creates a validation error with multiple errors.
func NewValidationErrors(errors []error) *ValidationError {
	return &ValidationError{
		errors: errors,
	}
}

// NewValidationError creates a validation error with just one error.
func NewValidationError(anerror error) *ValidationError {
	return &ValidationError{
		errors: []error{anerror},
	}
}

func (v *ValidationError) addError(errormessage string) {
	errormessage = strings.TrimSpace(errormessage)
	if errormessage == "" {
		return
	}
	newerror := errors.New(errormessage)
	v.errors = append(v.errors, newerror)
}

// ValidatePayment validtes if payment data is correct.
func (n *NewPayment) ValidatePayment() *ValidationError {
	var result ValidationError

	if n.Amount < 0.0 {
		result.addError("payment amount cannot be less than zero")
	}

	if n.Amount > 5000001.0 {
		result.addError("maximum payment amount allowed is $5000001")
	}

	if strings.TrimSpace(n.CompanyKey) == "" {
		result.addError("company key is mandatory")
	}

	if strings.TrimSpace(n.TxID) == "" {
		result.addError("a client transaction ID must be provided")
	}

	if len(result.errors) == 0 {
		return nil
	}

	return &result
}
