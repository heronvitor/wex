package entities

import (
	"time"
)

type Purchase struct {
	ID              string
	Description     string
	Amount          float64
	TransactionDate time.Time
}

type PurchaseInCurrency struct {
	ID              string
	Description     string
	Amount          float64
	TransactionDate time.Time
	CurrencyRate    float64
	ConvertedAmount float64
}
