package schemas

import (
	"strconv"
	"time"
)

type CreatePurchaseInput struct {
	TransactionDate Date    `json:"transaction_date" binding:"required"`
	Description     string  `json:"description" binding:"required"`
	Amount          float64 `json:"amount" binding:"required"`
}

type CreatePurchaseOutput struct {
	ID              string    `json:"id"`
	TransactionDate time.Time `json:"transaction_date"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
}

type GetPurchaseOutput struct {
	ID              string    `json:"id"`
	Description     string    `json:"description"`
	Amount          float64   `json:"amount"`
	TransactionDate time.Time `json:"transaction_date"`
	CurrencyRate    float64   `json:"currency_rate"`
	ConvertedAmount float64   `json:"converted_amount"`
}

type GetPurchaseInput struct {
	ID       string `form:"id" binding:"required"`
	Country  string `form:"country" binding:"required"`
	Currency string `form:"currency" binding:"required"`
}

type Date time.Time

func (dt *Date) UnmarshalJSON(data []byte) (err error) {
	str, err := strconv.Unquote(string(data))
	if err != nil {
		return
	}

	t, err := time.ParseInLocation(time.DateOnly, str, time.UTC)
	if err != nil {
		return
	}
	*dt = Date(t)
	return
}
