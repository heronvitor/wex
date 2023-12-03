package fiscaldata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter_encode(t *testing.T) {
	t.Run("", func(t *testing.T) {
		f := Filter{
			Field:     "record_date",
			Operation: FilterOperationEqual,
			Value:     "2023-09-30",
		}
		want := "record_date:eq:2023-09-30"

		got := f.encode()
		assert.Equal(t, got, want)
	})
}

func TestExchangeRatesOptions_encode(t *testing.T) {
	t.Run("page options", func(t *testing.T) {
		o := ExchangeRatesOptions{
			PageNumber: 12,
			PageSize:   5,
		}
		wantParams := map[string][]string{
			"format":       {"json"},
			"page[number]": {"12"},
			"page[size]":   {"5"},
		}

		gotParams := o.encode()
		assert.Equal(t, gotParams, wantParams)
	})

	t.Run("fields option", func(t *testing.T) {
		o := ExchangeRatesOptions{
			Fields: []string{"record_date", "country,currency", "country_currency_desc", "exchange_rate"},
		}

		wantParams := map[string][]string{
			"format": {"json"},
			"fields": {"record_date,country,currency,country_currency_desc,exchange_rate"},
		}

		gotParams := o.encode()
		assert.Equal(t, gotParams, wantParams)
	})

}
