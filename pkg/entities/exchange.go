package entities

import "time"

type ExchangeRate struct {
	RecordDate    string
	Country       string
	Currency      string
	ExchangeRate  int
	EffectiveDate string
}

type ExchangeRateUpdateInfo struct {
	LastAtemptDate time.Time
	RetryCount     int
	RetryDate      int
	Success        bool
}
