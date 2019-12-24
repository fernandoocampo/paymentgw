package domain

import (
	"time"

	"github.com/google/uuid"
)

// Payment contains required data for a payment transaction.
type Payment struct {
	ID         string
	TxID       string
	CompanyKey string
	Amount     float32
	Created    time.Time
}

// NewPayment creates a new payment with the given data. ID is generated internally.
func NewPayment(txID, companyKey string, amount float32) Payment {
	return Payment{
		ID:         uuid.New().String(),
		TxID:       txID,
		CompanyKey: companyKey,
		Amount:     amount,
		Created:    time.Now(),
	}
}
