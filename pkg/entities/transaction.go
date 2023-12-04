package entities

import (
	"time"

	"github.com/google/uuid"
)

type Purchase struct {
	UUID            uuid.UUID
	Description     string
	Amount          int
	TransactionDate time.Time
}

type PurchaseInCurrency struct {
	UUID            uuid.UUID
	Description     string
	Amount          int
	TransactionDate time.Time
	CurrencyRate    int
	ConvertedAmount int
}
