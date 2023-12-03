package fiscaldata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	URL    string
	Client *http.Client
}

func (c *Client) GetExchangeRates(options ExchangeRatesOptions) (exchangeRates ExchangeRatesResponse, err error) {
	params := url.Values(options.encode())

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
