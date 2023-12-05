package business

import (
	"errors"

	"github.com/google/uuid"
	"github.com/heronvitor/pkg/entities"
)

var (
	ErrCantConvertCurrency = errors.New("cannot be converted to the target currency")
	ErrPurchaseNotFound    = errors.New("purchase transaction not found")
	newUuid                = uuid.NewString
)

type PurchaseRepository interface {
	GetPurchaseByID(id string) (*entities.Purchase, error)
	SavePurchase(entities.Purchase) error
}

type PurchaseService struct {
	PurchaseRepository     PurchaseRepository
	ExchangeRateRepository ExchangeRateRepository
}

func (s PurchaseService) CreatePurchase(purchase entities.Purchase) (entities.Purchase, error) {
	purchase.ID = newUuid()

	err := s.PurchaseRepository.SavePurchase(purchase)
	if err != nil {
		return entities.Purchase{}, err
	}

	return purchase, err
}

func (s PurchaseService) GetPurchaseInCurrency(id string, currency string) (*entities.PurchaseInCurrency, error) {
	purchase, err := s.PurchaseRepository.GetPurchaseByID(id)
	if err != nil {
		return nil, err
	}

	if purchase == nil {
		return nil, ErrPurchaseNotFound
	}

	exchangeRate, err := s.ExchangeRateRepository.GetCurrencyRateUntil(currency, purchase.TransactionDate)
	if err != nil {
		return nil, err
	}

	if exchangeRate == nil {
		return nil, ErrCantConvertCurrency
	}

	return &entities.PurchaseInCurrency{
		ID:              purchase.ID,
		Description:     purchase.Description,
		Amount:          purchase.Amount,
		TransactionDate: purchase.TransactionDate,
		CurrencyRate:    exchangeRate.ExchangeRate,
		ConvertedAmount: purchase.Amount * exchangeRate.ExchangeRate,
	}, nil
}
