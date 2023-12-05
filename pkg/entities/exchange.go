package entities

import "time"

type ExchangeRate struct {
	RecordDate    string
	Country       string
	Currency      string
	ExchangeRate  float64
	EffectiveDate string
}

type ExchangeRateUpdateInfo struct {
	Time       time.Time
	RetryCount int
	RetryTime  time.Time
	Success    bool
}
