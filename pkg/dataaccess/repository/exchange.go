package repository

import (
	"database/sql"
	"time"

	"github.com/heronvitor/pkg/entities"
)

type ExangeRateRepository struct {
	DB *sql.DB
}

func (r ExangeRateRepository) SaveExchangeRates([]entities.ExchangeRate, entities.ExchangeRateUpdateInfo) error {
	return nil
}

func (r ExangeRateRepository) GetLastUpdateAttempt() (lastUpdateAttempt *entities.ExchangeRateUpdateInfo, err error) {
	return nil, nil
}

func (r ExangeRateRepository) GetCurrencyRateUntil(currency string, until time.Time) (*entities.ExchangeRate, error) {
	return nil, nil
}
