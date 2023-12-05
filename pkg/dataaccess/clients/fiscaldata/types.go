package fiscaldata

type ExchangeRatesResponse struct {
	ExchangeRates []ExchangeRate `json:"data"`
	Meta          struct {
		TotalPages int `json:"total-pages"`
	} `json:"meta"`
}

type ExchangeRate struct {
	RecordDate    string `json:"record_date"`
	Country       string `json:"country"`
	Currency      string `json:"currency"`
	ExchangeRate  string `json:"exchange_rate"`
	EffectiveDate string `json:"effective_date"`
}
