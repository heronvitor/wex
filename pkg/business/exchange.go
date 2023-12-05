package business

import (
	"log"
	"strconv"
	"time"

	"github.com/heronvitor/pkg/dataaccess/clients/fiscaldata"
	"github.com/heronvitor/pkg/entities"
)

const MaxUpdateRetries = 5

type ExchangeRateRepository interface {
	SaveExchangeRates([]entities.ExchangeRate, entities.ExchangeRateUpdateInfo) error
	GetLastUpdateAttempt() (lastUpdateAttempt *entities.ExchangeRateUpdateInfo, err error)
	GetCurrencyRateUntil(country, currency string, until time.Time) (*entities.ExchangeRate, error)
}

type ExchangeRatesClient interface {
	GetAllExchangeRates() (exchangeRates []fiscaldata.ExchangeRate, err error)
}

type ExchangeRatesService struct {
	ExchangeRateRepository
	ExchangeRatesClient
}

type UpdateOptions struct {
	Interval, RetryInterval time.Duration
	Now                     time.Time
}

func (ef ExchangeRatesService) Update(options UpdateOptions) (err error) {
	lastUpdateAttempt, err := ef.ExchangeRateRepository.GetLastUpdateAttempt()
	if err != nil {
		log.Printf("get latest update error: %s", err.Error())
		return
	}

	var updateInfo entities.ExchangeRateUpdateInfo

	if lastUpdateAttempt == nil {
		updateInfo = entities.ExchangeRateUpdateInfo{
			Time: options.Now,
		}
	} else if lastUpdateAttempt.Success {
		if options.Now.Sub(lastUpdateAttempt.Time) < options.Interval {
			log.Println("skiping update due to a recent update")
			return
		}
	} else {
		if options.Now.Sub(lastUpdateAttempt.RetryTime) < options.RetryInterval {
			log.Println("skiping update due to a recent attempt")
			return
		}

		updateInfo = *lastUpdateAttempt
		updateInfo.RetryTime = options.Now
		updateInfo.RetryCount += 1

	}
	exchangeRatesResp, err := ef.ExchangeRatesClient.GetAllExchangeRates()

	if err != nil {
		log.Printf("get exchange rates error: %s", err.Error())

		if err := ef.ExchangeRateRepository.SaveExchangeRates(nil, updateInfo); err != nil {
			log.Printf("update failed attempt error: %s", err.Error())
		}
		return
	}

	exchangeRates, err := convertRates(exchangeRatesResp)
	updateInfo.Success = true

	err = ef.ExchangeRateRepository.SaveExchangeRates(exchangeRates, updateInfo)
	if err != nil {
		log.Printf("save exchange rates error: %s", err.Error())
	}
	return
}

func convertRates(exchangeRates []fiscaldata.ExchangeRate) ([]entities.ExchangeRate, error) {
	converted := make([]entities.ExchangeRate, len(exchangeRates))

	rates := exchangeRates

	for i := range exchangeRates {
		recordDate, err := time.Parse(time.DateOnly, rates[i].RecordDate)
		if err != nil {
			return nil, err
		}

		effectiveDate, err := time.Parse(time.DateOnly, rates[i].EffectiveDate)
		if err != nil {
			return nil, err
		}
		exchangeRate, err := strconv.ParseFloat(rates[i].ExchangeRate, 64)
		if err != nil {
			return nil, err
		}

		converted = append(
			converted,
			entities.ExchangeRate{
				RecordDate:    recordDate,
				Country:       rates[i].Country,
				Currency:      rates[i].Currency,
				ExchangeRate:  exchangeRate,
				EffectiveDate: effectiveDate,
			},
		)
	}
	return converted, nil
}
