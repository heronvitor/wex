package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/heronvitor/pkg/entities"
)

type ExangeRateRepository struct {
	DB *sql.DB
}

func (r ExangeRateRepository) SaveExchangeRates(exchangeRates []entities.ExchangeRate, updateInfo entities.ExchangeRateUpdateInfo) error {
	query := `
		INSERT INTO exchange_rate (record_date, country, currency, exchange_rate, effective_date)
			VALUES ($1,$2,$3,$4,$5)
		ON CONFLICT(country, currency, record_date) DO UPDATE SET
			exchange_rate=EXCLUDED.exchange_rate,
			effective_date=EXCLUDED.effective_date
			`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return err
	}

	for _, exchangeRate := range exchangeRates {
		_, err = stmt.Exec(
			exchangeRate.RecordDate,
			exchangeRate.Country,
			exchangeRate.Currency,
			exchangeRate.ExchangeRate,
			exchangeRate.EffectiveDate,
		)
		if err != nil {
			fmt.Println(&exchangeRate)
			return err
		}

	}

	query = `
		INSERT INTO exchange_rate_update_info (time,retry_count,retry_time,success) 
			VALUES ($1,$2,$3,$4)
		ON CONFLICT(time) DO UPDATE SET
			retry_count=EXCLUDED.retry_count,
			retry_time=EXCLUDED.retry_time,
			success=EXCLUDED.success
	`
	_, err = r.DB.Exec(query, updateInfo.Time, updateInfo.RetryCount, updateInfo.RetryTime, updateInfo.Success)
	return err
}

func (r ExangeRateRepository) GetLastUpdateAttempt() (lastUpdateAttempt *entities.ExchangeRateUpdateInfo, err error) {
	lastUpdateAttempt = &entities.ExchangeRateUpdateInfo{}

	query := `
		SELECT time,retry_count,retry_time,success
		FROM exchange_rate_update_info
		ORDER BY time DESC
	`
	err = r.DB.QueryRow(query).Scan(
		&lastUpdateAttempt.Time,
		&lastUpdateAttempt.RetryCount,
		&lastUpdateAttempt.RetryTime,
		&lastUpdateAttempt.Success,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	lastUpdateAttempt.Time = lastUpdateAttempt.Time.UTC()
	lastUpdateAttempt.RetryTime = lastUpdateAttempt.RetryTime.UTC()
	return lastUpdateAttempt, nil
}

func (r ExangeRateRepository) GetCurrencyRateUntil(country, currency string, until time.Time) (*entities.ExchangeRate, error) {
	exchangeRate := &entities.ExchangeRate{}
	query := `
		SELECT record_date, country, currency, exchange_rate, effective_date 
			FROM exchange_rate 
			WHERE country=$1 and currency=$2 and record_date<=$3
			ORDER BY record_date DESC
	`
	row := r.DB.QueryRow(query, country, currency, until)

	err := row.Scan(
		&exchangeRate.RecordDate,
		&exchangeRate.Country,
		&exchangeRate.Currency,
		&exchangeRate.ExchangeRate,
		&exchangeRate.EffectiveDate,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	exchangeRate.EffectiveDate = exchangeRate.EffectiveDate.UTC()
	exchangeRate.RecordDate = exchangeRate.RecordDate.UTC()
	return exchangeRate, nil
}
