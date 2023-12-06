package business

import (
	"errors"
	"math"

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

func (s PurchaseService) GetPurchaseInCurrency(id, country, currency string) (*entities.PurchaseInCurrency, error) {
	purchase, err := s.PurchaseRepository.GetPurchaseByID(id)
	if err != nil {
		return nil, err
	}

	if purchase == nil {
		return nil, ErrPurchaseNotFound
	}

	exchangeRate, err := s.ExchangeRateRepository.GetCurrencyRateUntil(country, currency, purchase.TransactionDate)
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
		ConvertedAmount: roundFloat(purchase.Amount*exchangeRate.ExchangeRate, 2),
	}, nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
