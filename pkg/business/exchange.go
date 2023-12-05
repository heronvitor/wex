package business

import (
	"time"

	"github.com/heronvitor/pkg/dataaccess/clients/fiscaldata"
	"github.com/heronvitor/pkg/entities"
)

const MaxUpdateRetries = 5

type ExchangeRateRepository interface {
	SaveExchangeRates([]entities.ExchangeRate, entities.ExchangeRateUpdateInfo) error
	GetLastUpdateAttempt() (lastUpdateAttempt *entities.ExchangeRateUpdateInfo, err error)
	GetCurrencyRateUntil(currency string, until time.Time) (*entities.ExchangeRate, error)
}

type ExchangeRatesClient interface {
	GetAllExchangeRates() (exchangeRates fiscaldata.ExchangeRatesResponse, err error)
}

type ExchangeRatesService struct {
	ExchangeRateRepository
	ExchangeRatesClient
}

type UpdateOptions struct {
	UpdateSince             *time.Time
	Interval, RetryInterval time.Duration
	Now                     time.Time
}

func (ef ExchangeRatesService) Update(options UpdateOptions) (err error) {
	lastUpdateAttempt, err := ef.ExchangeRateRepository.GetLastUpdateAttempt()
	if err != nil {
		return
	}

	var updateInfo entities.ExchangeRateUpdateInfo

	if lastUpdateAttempt != nil && !lastUpdateAttempt.Success {
		if options.Now.Sub(lastUpdateAttempt.RetryTime) < options.RetryInterval {
			return
		}

		updateInfo = *lastUpdateAttempt
		updateInfo.RetryTime = options.Now
		updateInfo.RetryCount += 1

	} else {
		updateInfo = entities.ExchangeRateUpdateInfo{
			Time: options.Now,
		}
	}

	exchangeRatesResp, err := ef.ExchangeRatesClient.GetAllExchangeRates()
	if err != nil {
		err = ef.ExchangeRateRepository.SaveExchangeRates(nil, updateInfo)
		return
	}

	exchangeRates := convertRates(exchangeRatesResp)
	updateInfo.Success = true

	err = ef.ExchangeRateRepository.SaveExchangeRates(exchangeRates, updateInfo)
	return
}

func convertRates(exchangeRatesRes fiscaldata.ExchangeRatesResponse) []entities.ExchangeRate {
	converted := make([]entities.ExchangeRate, len(exchangeRatesRes.ExchangeRates))

	rates := exchangeRatesRes.ExchangeRates
	for i := range exchangeRatesRes.ExchangeRates {
		converted = append(
			converted,
			entities.ExchangeRate{
				RecordDate:    rates[i].RecordDate,
				Country:       rates[i].Country,
				Currency:      rates[i].Currency,
				ExchangeRate:  0, //rates[i].ExchangeRate,
				EffectiveDate: rates[i].EffectiveDate,
			},
		)
	}
	return nil
}
