package fiscaldata

type ExchangeRatesResponse struct {
	ExchangeRates []ExchangeRate `json:"data"`
	TotalPages    int            `json:"total_pages"`
}

type ExchangeRate struct {
	RecordDate    string `json:"record_date"`
	Country       string `json:"country"`
	Currency      string `json:"currency"`
	ExchangeRate  string `json:"exchange_rate"`
	EffectiveDate string `json:"effective_date"`
}
