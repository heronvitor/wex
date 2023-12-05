package rest

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	mocks "github.com/heronvitor/mocks/pkg/presentation/rest"
	"github.com/heronvitor/pkg/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPurchaseHandler_CreatePurchase(t *testing.T) {
	t.Run("should validate input", func(t *testing.T) {
		wantBody := `{"error":"Key: 'CreatePurchaseInput.Amount' Error:Field validation for 'Amount' failed on the 'required' tag"}`
		wantStatus := 400

		handler := PurchaseHandler{}
		req, err := http.NewRequest("POST", "", strings.NewReader(`{"description":"purchase desc","transaction_date":"2023-04-20"}`))
		assert.NoError(t, err)

		rr := callGinHandler(t, handler.CreatePurchase, req)
		assert.Equal(t, wantStatus, rr.Code)
		assert.Equal(t, wantBody, rr.Body.String())
	})

	t.Run("should return call create", func(t *testing.T) {
		purchaseService := &mocks.PurchaseService{}

		handler := PurchaseHandler{
			PurchaseService: purchaseService,
		}

		purchaseService.On("CreatePurchase", entities.Purchase{
			Description:     "purchase desc",
			Amount:          24.56,
			TransactionDate: time.Date(2023, time.April, 20, 0, 0, 0, 0, time.UTC),
		}).Return(entities.Purchase{}, errors.New("service error"))

		req, err := http.NewRequest("POST", "", strings.NewReader(`{"description":"purchase desc","amount":24.56,"transaction_date":"2023-04-20"}`))
		assert.NoError(t, err)

		_ = callGinHandler(t, handler.CreatePurchase, req)
		purchaseService.AssertExpectations(t)
	})

	t.Run("should return internal error error", func(t *testing.T) {
		purchaseService := &mocks.PurchaseService{}

		handler := PurchaseHandler{
			PurchaseService: purchaseService,
		}

		wantBody := `{"error":"internal error"}`
		wantStatus := 500

		purchaseService.On("CreatePurchase", mock.Anything).
			Return(entities.Purchase{}, errors.New("service error"))

		req, err := http.NewRequest("POST", "", strings.NewReader(`{"description":"purchase desc","amount":24.56,"transaction_date":"2023-04-20"}`))
		assert.NoError(t, err)

		rr := callGinHandler(t, handler.CreatePurchase, req)
		assert.Equal(t, wantStatus, rr.Code)
		assert.Equal(t, wantBody, rr.Body.String())
	})

	t.Run("should return created", func(t *testing.T) {
		purchaseService := &mocks.PurchaseService{}

		handler := PurchaseHandler{
			PurchaseService: purchaseService,
		}

		wantBody := `{"id":"bc3ede90-ecae-424a-97d4-2ca9556fd5e7","transaction_date":"2023-04-20T00:00:00Z","description":"purchase desc","amount":24.56}`

		purchaseService.On("CreatePurchase", mock.Anything).
			Return(
				entities.Purchase{
					ID:              "bc3ede90-ecae-424a-97d4-2ca9556fd5e7",
					Description:     "purchase desc",
					Amount:          24.56,
					TransactionDate: time.Date(2023, time.April, 20, 0, 0, 0, 0, time.UTC)},
				nil,
			)

		req, err := http.NewRequest("POST", "", strings.NewReader(`{"description":"purchase desc","amount":24.56,"transaction_date":"2023-04-20"}`))
		assert.NoError(t, err)

		rr := callGinHandler(t, handler.CreatePurchase, req)
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, wantBody, rr.Body.String())
	})
}

func TestPurchaseHandler_GetPurchase(t *testing.T) {
	t.Run("should require params", func(t *testing.T) {
		wantBody := `{"error":"Key: 'GetPurchaseInput.Currency' Error:Field validation for 'Currency' failed on the 'required' tag"}`
		wantStatus := 400

		handler := PurchaseHandler{}
		req, err := http.NewRequest("GET", "?id=bc3ede90-ecae-424a-97d4-2ca9556fd5e7", strings.NewReader(`{"description":"purchase desc","transaction_date":"2023-04-20"}`))
		assert.NoError(t, err)

		rr := callGinHandler(t, handler.GetPurchase, req)
		assert.Equal(t, wantStatus, rr.Code)
		assert.Equal(t, wantBody, rr.Body.String())
	})

	t.Run("should call get purchase", func(t *testing.T) {
		purchaseService := &mocks.PurchaseService{}

		handler := PurchaseHandler{
			PurchaseService: purchaseService,
		}
		req, err := http.NewRequest("GET", "?id=bc3ede90-ecae-424a-97d4-2ca9556fd5e7&currency=real", strings.NewReader(`{"description":"purchase desc","transaction_date":"2023-04-20"}`))
		assert.NoError(t, err)

		purchaseService.On("GetPurchaseInCurrency", "bc3ede90-ecae-424a-97d4-2ca9556fd5e7", "real").
			Return(
				nil,
				nil,
			)

		_ = callGinHandler(t, handler.GetPurchase, req)
		purchaseService.AssertExpectations(t)
	})

	t.Run("should return not found", func(t *testing.T) {
		purchaseService := &mocks.PurchaseService{}

		wantBody := `{"error":"not found"}`
		wantStatus := 404

		handler := PurchaseHandler{
			PurchaseService: purchaseService,
		}
		req, err := http.NewRequest("GET", "?id=bc3ede90-ecae-424a-97d4-2ca9556fd5e7&currency=real", strings.NewReader(`{"description":"purchase desc","transaction_date":"2023-04-20"}`))
		assert.NoError(t, err)

		purchaseService.On("GetPurchaseInCurrency", mock.Anything, mock.Anything).
			Return(
				nil,
				nil,
			)

		rr := callGinHandler(t, handler.GetPurchase, req)
		purchaseService.AssertExpectations(t)
		assert.Equal(t, wantStatus, rr.Code)
		assert.Equal(t, wantBody, rr.Body.String())
	})

	t.Run("should return not found", func(t *testing.T) {
		purchaseService := &mocks.PurchaseService{}

		wantBody := `{"error":"internal error"}`
		wantStatus := 500

		handler := PurchaseHandler{
			PurchaseService: purchaseService,
		}
		req, err := http.NewRequest("GET", "?id=bc3ede90-ecae-424a-97d4-2ca9556fd5e7&currency=real", strings.NewReader(`{"description":"purchase desc","transaction_date":"2023-04-20"}`))
		assert.NoError(t, err)

		purchaseService.On("GetPurchaseInCurrency", mock.Anything, mock.Anything).
			Return(
				nil,
				errors.New("service error"),
			)

		rr := callGinHandler(t, handler.GetPurchase, req)
		purchaseService.AssertExpectations(t)
		assert.Equal(t, wantStatus, rr.Code)
		assert.Equal(t, wantBody, rr.Body.String())
	})

	t.Run("should return purchase", func(t *testing.T) {
		purchaseService := &mocks.PurchaseService{}

		wantBody := `{"id":"da7cd6e6-1362-4761-af8a-b829b3ea7d60","description":"description","amount":10,"transaction_date":"2020-05-04T01:00:00Z","currency_rate":53,"converted_amount":530}`
		wantStatus := 200

		handler := PurchaseHandler{
			PurchaseService: purchaseService,
		}
		req, err := http.NewRequest("GET", "?id=bc3ede90-ecae-424a-97d4-2ca9556fd5e7&currency=real", strings.NewReader(`{"description":"purchase desc","transaction_date":"2023-04-20"}`))
		assert.NoError(t, err)

		purchaseService.On("GetPurchaseInCurrency", mock.Anything, mock.Anything).
			Return(
				&entities.PurchaseInCurrency{
					ID:              "da7cd6e6-1362-4761-af8a-b829b3ea7d60",
					Description:     "description",
					Amount:          10,
					ConvertedAmount: 530,
					CurrencyRate:    53,
					TransactionDate: time.Date(2020, 5, 4, 1, 0, 0, 0, time.UTC),
				},
				nil,
			)

		rr := callGinHandler(t, handler.GetPurchase, req)
		purchaseService.AssertExpectations(t)
		assert.Equal(t, wantStatus, rr.Code)
		assert.Equal(t, wantBody, rr.Body.String())
	})
}

func callGinHandler(t *testing.T, handler gin.HandlerFunc, request *http.Request) *httptest.ResponseRecorder {
	t.Helper()
	rr := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(rr)
	r.Use(handler)
	ctx.Request = request

	r.ServeHTTP(rr, ctx.Request)
	return rr
}
