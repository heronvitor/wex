package business

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	mocks "github.com/heronvitor/mocks/pkg/business"

	"github.com/heronvitor/pkg/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransactionService_SavePurchase(t *testing.T) {
	newUuid = func() string { return "da7cd6e6-1362-4761-af8a-b829b3ea7d60" }
	defer func() {
		newUuid = uuid.NewString
	}()

	t.Run("should set id and call save", func(t *testing.T) {
		transactionRepository := &mocks.TransactionRepository{}

		service := TransactionService{
			TransactionRepository: transactionRepository,
		}

		transactionRepository.On("SavePurchase", mock.Anything).
			Return(errors.New("save error"))

		_, gotErr := service.SavePurchase(entities.Purchase{
			Description:     "description",
			Amount:          10,
			TransactionDate: time.Date(2020, 5, 4, 1, 0, 0, 0, time.UTC),
		})

		transactionRepository.AssertExpectations(t)
		assert.Equal(t, gotErr, errors.New("save error"))
	})

	t.Run("should set id and call save", func(t *testing.T) {
		transactionRepository := &mocks.TransactionRepository{}
		wantPurchase := entities.Purchase{
			ID:              "da7cd6e6-1362-4761-af8a-b829b3ea7d60",
			Description:     "description",
			Amount:          10,
			TransactionDate: time.Date(2020, 5, 4, 1, 0, 0, 0, time.UTC),
		}

		service := TransactionService{
			TransactionRepository: transactionRepository,
		}

		transactionRepository.On("SavePurchase", wantPurchase).
			Return(nil)

		gotPurchase, gotErr := service.SavePurchase(entities.Purchase{
			Description:     "description",
			Amount:          10,
			TransactionDate: time.Date(2020, 5, 4, 1, 0, 0, 0, time.UTC),
		})

		transactionRepository.AssertExpectations(t)
		assert.Equal(t, wantPurchase, gotPurchase)
		assert.NoError(t, gotErr)
	})
}

func TestTransactionService_GetPurchaseInCurrency(t *testing.T) {
	id := "da7cd6e6-1362-4761-af8a-b829b3ea7d60"

	t.Run("should pass error", func(t *testing.T) {
		transactionRepository := &mocks.TransactionRepository{}

		service := TransactionService{
			TransactionRepository: transactionRepository,
		}

		transactionRepository.On("GetPurchaseByID", id).
			Return(
				nil,
				errors.New("random error"),
			)

		gotPurchase, gotErr := service.GetPurchaseInCurrency(id, "real")

		assert.Equal(t, gotErr, errors.New("random error"))
		assert.Nil(t, gotPurchase)
	})

	t.Run("should return not found error", func(t *testing.T) {
		transactionRepository := &mocks.TransactionRepository{}

		service := TransactionService{
			TransactionRepository: transactionRepository,
		}

		transactionRepository.On("GetPurchaseByID", id).
			Return(nil, nil)

		gotPurchase, gotErr := service.GetPurchaseInCurrency(id, "real")

		assert.Equal(t, gotErr, ErrPurchaseNotFound)
		assert.Nil(t, gotPurchase)

	})

	t.Run("should return get exchange error", func(t *testing.T) {
		transactionRepository := &mocks.TransactionRepository{}
		exchangeRateRepository := &mocks.ExchangeRateRepository{}

		service := TransactionService{
			TransactionRepository:  transactionRepository,
			ExchangeRateRepository: exchangeRateRepository,
		}

		transactionRepository.On("GetPurchaseByID", mock.Anything).
			Return(
				&entities.Purchase{},
				nil,
			)

		exchangeRateRepository.On("GetCurrencyRateUntil", mock.Anything, mock.Anything).
			Return(
				&entities.ExchangeRate{},
				errors.New("error"),
			)

		gotPurchase, gotErr := service.GetPurchaseInCurrency(id, "real")

		assert.Equal(t, gotErr, errors.New("error"))
		assert.Nil(t, gotPurchase)
	})

	t.Run("should return cant convert error", func(t *testing.T) {
		transactionRepository := &mocks.TransactionRepository{}
		exchangeRateRepository := &mocks.ExchangeRateRepository{}

		service := TransactionService{
			TransactionRepository:  transactionRepository,
			ExchangeRateRepository: exchangeRateRepository,
		}

		transactionRepository.On("GetPurchaseByID", mock.Anything).
			Return(
				&entities.Purchase{},
				nil,
			)

		exchangeRateRepository.On("GetCurrencyRateUntil", mock.Anything, mock.Anything).
			Return(
				nil,
				nil,
			)

		gotPurchase, gotErr := service.GetPurchaseInCurrency(id, "real")

		assert.Equal(t, gotErr, errors.New("cannot be converted to the target currency"))
		assert.Nil(t, gotPurchase)
	})

	t.Run("should set id and call save", func(t *testing.T) {
		transactionRepository := &mocks.TransactionRepository{}
		exchangeRateRepository := &mocks.ExchangeRateRepository{}

		service := TransactionService{
			TransactionRepository:  transactionRepository,
			ExchangeRateRepository: exchangeRateRepository,
		}

		wantPurchase := &entities.PurchaseInCurrency{
			ID:              "da7cd6e6-1362-4761-af8a-b829b3ea7d60",
			Description:     "description",
			Amount:          10,
			ConvertedAmount: 530,
			CurrencyRate:    53,
			TransactionDate: time.Date(2020, 5, 4, 1, 0, 0, 0, time.UTC),
		}

		transactionRepository.On("GetPurchaseByID", id).
			Return(
				&entities.Purchase{
					ID:              id,
					Description:     "description",
					Amount:          10,
					TransactionDate: time.Date(2020, 5, 4, 1, 0, 0, 0, time.UTC),
				},
				nil,
			)

		exchangeRateRepository.On("GetCurrencyRateUntil", "real", time.Date(2020, 5, 4, 1, 0, 0, 0, time.UTC)).
			Return(
				&entities.ExchangeRate{
					Country:      "Brazil",
					Currency:     "real",
					ExchangeRate: 53,
				},
				nil,
			)

		gotPurchase, gotErr := service.GetPurchaseInCurrency(id, "real")

		assert.Equal(t, wantPurchase, gotPurchase)
		assert.NoError(t, gotErr)

		transactionRepository.AssertExpectations(t)
		exchangeRateRepository.AssertExpectations(t)
	})
}
