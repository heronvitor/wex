package repository

import (
	"testing"
	"time"

	"github.com/heronvitor/pkg/entities"
	"github.com/stretchr/testify/assert"
)

func TestExangeRateRepository_GetCurrencyRateUntil(t *testing.T) {

	t.Run("should echange rate for currency until", func(t *testing.T) {
		db := createDB(t)
		defer db.Close()

		// prepare db
		query := "INSERT INTO exchange_rate (record_date, country, currency, exchange_rate, effective_date) VALUES ($1,$2,$3,$4,$5)"
		_, err := db.Exec(query, "2022-01-02", "Brazil", "real", "3.5", "2022-01-02")
		assert.NoError(t, err)
		_, err = db.Exec(query, "2023-05-02", "Brazil", "real", "5.3", "2023-05-02")
		assert.NoError(t, err)
		_, err = db.Exec(query, "2024-05-02", "Brazil", "real", "6.3", "2024-01-02")
		assert.NoError(t, err)
		_, err = db.Exec(query, "2022-06-02", "Argentina", "peso", "10000", "2023-06-02")
		assert.NoError(t, err)

		// actual test
		wantExchangeRate := &entities.ExchangeRate{
			RecordDate:    time.Date(2023, 5, 2, 0, 0, 0, 0, time.UTC),
			Country:       "Brazil",
			Currency:      "real",
			ExchangeRate:  5.3,
			EffectiveDate: time.Date(2023, 5, 2, 0, 0, 0, 0, time.UTC),
		}

		repository := ExangeRateRepository{DB: db}

		gotExchangeRate, goErr := repository.GetCurrencyRateUntil("real", time.Date(2023, 6, 2, 0, 0, 0, 0, time.UTC))

		assert.NoError(t, goErr)
		// assert.Equal(t, wantExchangeRate, gotExchangeRate)
		assert.Equal(t, wantExchangeRate.EffectiveDate, gotExchangeRate.EffectiveDate)
	})
}

func TestExangeRateRepository_SaveExchangeRates(t *testing.T) {
	t.Run("should save update info", func(t *testing.T) {
		db := createDB(t)
		defer db.Close()

		repository := ExangeRateRepository{DB: db}

		goErr := repository.SaveExchangeRates(
			[]entities.ExchangeRate{},
			entities.ExchangeRateUpdateInfo{
				Time:       time.Date(2023, 03, 17, 20, 38, 16, 0, time.UTC),
				RetryCount: 5,
				RetryTime:  time.Date(2023, 03, 17, 23, 38, 16, 0, time.UTC),
				Success:    true,
			},
		)

		assert.Nil(t, goErr)

		// check db
		fields := make([]string, 4, 4)

		assert.NoError(t, db.QueryRow("SELECT * FROM exchange_rate_update_info").
			Scan(&fields[0], &fields[1], &fields[2], &fields[3]))

		wantFields := []string{"2023-03-17T20:38:16Z", "5", "2023-03-17T23:38:16Z", "true"}

		assert.Equal(t, wantFields, fields)
	})

	t.Run("should create exchange_rates", func(t *testing.T) {
		db := createDB(t)
		defer db.Close()

		repository := ExangeRateRepository{DB: db}

		goErr := repository.SaveExchangeRates(
			[]entities.ExchangeRate{
				{
					RecordDate:    time.Date(2023, 5, 2, 0, 0, 0, 0, time.UTC),
					Country:       "Brazil",
					Currency:      "real",
					ExchangeRate:  5.3,
					EffectiveDate: time.Date(2023, 5, 2, 0, 0, 0, 0, time.UTC),
				},
				{
					RecordDate:    time.Date(2023, 5, 2, 0, 0, 0, 0, time.UTC),
					Country:       "Argentina",
					Currency:      "peso",
					ExchangeRate:  100000,
					EffectiveDate: time.Date(2023, 5, 2, 0, 0, 0, 0, time.UTC),
				},
			},
			entities.ExchangeRateUpdateInfo{},
		)

		assert.NoError(t, goErr)

		// check db

		rows, err := db.Query("SELECT * FROM exchange_rate")
		assert.NoError(t, err)

		defer rows.Close()

		results := [][]string{}

		for rows.Next() {
			fields := make([]string, 5, 5)
			err = rows.Scan(&fields[0], &fields[1], &fields[2], &fields[3], &fields[4])

			results = append(results, fields)
			assert.NoError(t, err)
		}

		wantResults := [][]string{
			{"2023-05-02T00:00:00Z", "Brazil", "real", "5.30", "2023-05-02T00:00:00Z"},
			{"2023-05-02T00:00:00Z", "Argentina", "peso", "100000.00", "2023-05-02T00:00:00Z"},
		}
		assert.Equal(t, wantResults, results)
	})
}

func TestExangeRateRepository_GetLastUpdateAttempt(t *testing.T) {
	t.Run("should return nothing", func(t *testing.T) {
		db := createDB(t)
		defer db.Close()

		repository := ExangeRateRepository{DB: db}

		gotLastUpdateAttempt, gotErr := repository.GetLastUpdateAttempt()
		assert.NoError(t, gotErr)
		assert.Nil(t, gotLastUpdateAttempt)
	})

	t.Run("should return last update attempt", func(t *testing.T) {
		db := createDB(t)
		defer db.Close()

		// prepare db
		query := "INSERT INTO exchange_rate_update_info (time,retry_count,retry_time,success) VALUES ($1,$2,$3,$4)"
		_, err := db.Exec(query, "2022-03-17 20:38:16", 2, "2022-03-17 23:38:16", false)
		assert.NoError(t, err)
		_, err = db.Exec(query, "2023-03-17 20:38:16", 5, "2023-03-17 23:38:16", true)

		// test
		wantLastUpdateAttempt := &entities.ExchangeRateUpdateInfo{
			Time:       time.Date(2023, 03, 17, 20, 38, 16, 0, time.UTC),
			RetryCount: 5,
			RetryTime:  time.Date(2023, 03, 17, 23, 38, 16, 0, time.UTC),
			Success:    true,
		}

		repository := ExangeRateRepository{DB: db}

		gotLastUpdateAttempt, gotErr := repository.GetLastUpdateAttempt()
		assert.NoError(t, gotErr)
		assert.Equal(t, wantLastUpdateAttempt, gotLastUpdateAttempt)
	})
}
