package business

import (
	"errors"
	"testing"

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

		exchangeRateRepository.On("GetLastUpdateAttempt").
			Return(&entities.ExchangeRateUpdateInfo{}, errors.New("get last attempt error"))

		gotErr := service.Update(UpdateOptions{})

		exchangeRateRepository.AssertExpectations(t)
		assert.Equal(t, gotErr, errors.New("get last attempt error"))
	})

}
