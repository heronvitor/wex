package fiscaldata

import (
	"fmt"
	"strconv"
	"strings"
)

type FilterOperation string

const (
	FilterOperationLess         FilterOperation = "lt"
	FilterOperationLessEqual    FilterOperation = "lte"
	FilterOperationGreater      FilterOperation = "gt"
	FilterOperationGreaterEqual FilterOperation = "gte"
	FilterOperationEqual        FilterOperation = "eq"
	FilterOperationIn           FilterOperation = "in"
)

type Filter struct {
	Field     string
	Operation FilterOperation
	Value     string
}

func (f Filter) encode() string {
	return fmt.Sprintf("%s:%s:%s", f.Field, f.Operation, f.Value)
}

type ExchangeRatesOptions struct {
	Fields     []string
	PageSize   int
	PageNumber int
	SortFields []string
	Filters    []Filter
}

func (o ExchangeRatesOptions) encode() (params map[string][]string) {
	params = map[string][]string{
		"format": {"json"},
	}

	if len(o.Fields) != 0 {
		params["fields"] = []string{strings.Join(o.Fields, ",")}
	}
	if o.PageSize != 0 {
		params["page[size]"] = []string{strconv.Itoa(o.PageSize)}
	}
	if o.PageNumber != 0 {
		params["page[number]"] = []string{strconv.Itoa(o.PageNumber)}
	}
	if o.PageNumber != 0 {
		params["page[number]"] = []string{strconv.Itoa(o.PageNumber)}
	}
	if len(o.SortFields) != 0 {
		params["sort"] = []string{strings.Join(o.SortFields, ",")}
	}

	if len(o.Filters) != 0 {
		filters := []string{}
		for _, filter := range o.Filters {
			filters = append(filters, filter.encode())
		}

		strings.Join(filters, ",")
	}
	return params
}

type ExchangeRatesResponse struct {
	ExchangeRates []ExchangeRate `json:"data"`
}

type ExchangeRate struct {
	RecordDate    string `json:"record_date"`
	Country       string `json:"country"`
	Currency      string `json:"currency"`
	ExchangeRate  string `json:"exchange_rate"`
	EffectiveDate string `json:"effective_date"`
}
