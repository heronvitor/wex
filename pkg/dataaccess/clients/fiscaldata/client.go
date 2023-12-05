package fiscaldata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	URL    string
	Client *http.Client
}

func (c *Client) GetAllExchangeRates() (exchangeRates []ExchangeRate, err error) {
	pageSize := 1000
	pageNumber := 1
	var res ExchangeRatesResponse

	for {
		res, err = c.GetExchangeRates(pageSize, pageNumber)
		if err != nil {
			return nil, err
		}
		exchangeRates = append(exchangeRates, res.ExchangeRates...)

		if pageNumber == res.TotalPages {
			return
		}
		pageNumber++

	}
}

func (c *Client) GetExchangeRates(pageSize, pageNumber int) (exchangeRates ExchangeRatesResponse, err error) {
	params := url.Values{
		"format": {"json"},
		"fields": {"record_date,country,currency,exchange_rate,effective_date"},
	}

	if pageSize != 0 {
		params["page[size]"] = []string{strconv.Itoa(pageSize)}
	}
	if pageNumber != 0 {
		params["page[number]"] = []string{strconv.Itoa(pageNumber)}
	}

	url := fmt.Sprintf(
		"%s/services/api/fiscal_service/v1/accounting/od/rates_of_exchange?%s",
		c.URL,
		params.Encode(),
	)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return
	}

	err = json.NewDecoder(res.Body).Decode(&exchangeRates)
	if err != nil {
		return
	}
	return
}
