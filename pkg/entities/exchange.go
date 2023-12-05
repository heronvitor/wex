package entities

import "time"

type ExchangeRate struct {
	RecordDate    time.Time
	Country       string
	Currency      string
	ExchangeRate  float64
	EffectiveDate time.Time
}

type ExchangeRateUpdateInfo struct {
	Time       time.Time
	RetryCount int
	RetryTime  time.Time
	Success    bool
}
