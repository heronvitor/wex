package fiscaldata

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	_ "embed"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/exchange_response.json
var exchange_response string

func TestClient_GetExchangeRates(t *testing.T) {
	t.Run("should send url params response", func(t *testing.T) {
		wantPath := "/services/api/fiscal_service/v1/accounting/od/rates_of_exchange"
		wantQuery := url.Values{
			"fields":       []string{"record_date,country,currency,exchange_rate,effective_date"},
			"format":       []string{"json"},
			"page[number]": []string{"4"},
			"page[size]":   []string{"10"},
		}

		var gotPath string
		var gotQuery url.Values
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			gotPath = req.URL.Path
			gotQuery = req.URL.Query()

			rw.WriteHeader(200)
			rw.Write([]byte("{}"))
		}))
		defer server.Close()

		client := &Client{Client: http.DefaultClient, URL: server.URL}
		client.GetExchangeRates(10, 4)

		assert.Equal(t, wantPath, gotPath)
		assert.Equal(t, wantQuery, gotQuery)
	})

	t.Run("should parse response", func(t *testing.T) {
		wantExchangeRates := ExchangeRatesResponse{
			ExchangeRates: []ExchangeRate{
				{
					RecordDate:    "2001-03-31",
					Country:       "Malawi",
					Currency:      "Kwacha",
					ExchangeRate:  "79.75",
					EffectiveDate: "2001-03-31",
				},
				{
					RecordDate:    "2001-03-31",
					Country:       "Malaysia",
					Currency:      "Ringgit",
					ExchangeRate:  "3.8",
					EffectiveDate: "2001-03-31",
				},
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(200)
			rw.Write([]byte(exchange_response))
		}))
		defer server.Close()

		client := &Client{Client: http.DefaultClient, URL: server.URL}
		gotExchangeRates, gotErr := client.GetExchangeRates(0, 0)

		assert.Equal(t, wantExchangeRates, gotExchangeRates)
		assert.NoError(t, gotErr)
	})
}
