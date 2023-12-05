package entities

import (
	"time"
)

type Purchase struct {
	UUID            string
	Description     string
	Amount          int
	TransactionDate time.Time
}

type PurchaseInCurrency struct {
	UUID            string
	Description     string
	Amount          int
	TransactionDate time.Time
	CurrencyRate    int
	ConvertedAmount int
}
