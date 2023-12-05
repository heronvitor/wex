package business

import (
	"errors"
	"testing"
	"time"

	mocks "github.com/heronvitor/mocks/pkg/business"
	"github.com/heronvitor/pkg/entities"
	"github.com/stretchr/testify/assert"
)

func TestExchangeRatesService_Update(t *testing.T) {

	t.Run("should register update atempt and return on error", func(t *testing.T) {
		exchangeRateRepository := &mocks.ExchangeRateRepository{}

		service := ExchangeRatesService{
			ExchangeRateRepository: exchangeRateRepository,
		}

		wantErr := errors.New("get last attempt error")

		exchangeRateRepository.On("GetLastUpdateAttempt").
			Return(&entities.ExchangeRateUpdateInfo{}, errors.New("get last attempt error"))

		gotErr := service.Update(UpdateOptions{})

		exchangeRateRepository.AssertExpectations(t)
		assert.Equal(t, gotErr, wantErr)
	})

	t.Run("should skip when updated recent", func(t *testing.T) {
		exchangeRateRepository := &mocks.ExchangeRateRepository{}

		service := ExchangeRatesService{
			ExchangeRateRepository: exchangeRateRepository,
		}

		exchangeRateRepository.On("GetLastUpdateAttempt").
			Return(
				&entities.ExchangeRateUpdateInfo{
					Time:    time.Date(2023, 4, 30, 0, 0, 0, 0, time.UTC),
					Success: true,
				},
				nil,
			)

		gotErr := service.Update(
			UpdateOptions{
				Interval: 24 * time.Hour,
				Now:      time.Date(2023, 4, 30, 0, 0, 0, 0, time.UTC),
			},
		)

		exchangeRateRepository.AssertExpectations(t)
		assert.Nil(t, gotErr)
	})

	t.Run("should skip when retried recent", func(t *testing.T) {
		exchangeRateRepository := &mocks.ExchangeRateRepository{}

		service := ExchangeRatesService{
			ExchangeRateRepository: exchangeRateRepository,
		}

		exchangeRateRepository.On("GetLastUpdateAttempt").
			Return(
				&entities.ExchangeRateUpdateInfo{
					Time:      time.Date(2023, 4, 29, 0, 0, 0, 0, time.UTC),
					RetryTime: time.Date(2023, 4, 30, 0, 0, 0, 0, time.UTC),
					Success:   false,
				},
				nil,
			)

		gotErr := service.Update(
			UpdateOptions{
				RetryInterval: 1 * time.Hour,
				Now:           time.Date(2023, 4, 30, 0, 0, 0, 0, time.UTC),
			},
		)

		exchangeRateRepository.AssertExpectations(t)
		assert.Nil(t, gotErr)
	})
}
