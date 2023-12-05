package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/heronvitor/pkg/business"
	"github.com/heronvitor/pkg/entities"
	"github.com/heronvitor/pkg/presentation/rest/schemas"
)

type PurchaseService interface {
	CreatePurchase(purchase entities.Purchase) (entities.Purchase, error)
	GetPurchaseInCurrency(id string, currency string) (*entities.PurchaseInCurrency, error)
}

type PurchaseHandler struct {
	PurchaseService PurchaseService
}

// @Produce      json
// @Param        params  body  schemas.CreatePurchaseInput  true  "purchase"
// @Success      200  {object}  schemas.CreatePurchaseOutput
// @Failure      400  {object}  Error
// @Failure      500  {object}  Error
// @Router       /accounts [post]
func (h *PurchaseHandler) CreatePurchase(c *gin.Context) {
	input := schemas.CreatePurchaseInput{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, schemas.Error{Error: err.Error()})
		return
	}

	purchase := entities.Purchase{
		Description:     input.Description,
		Amount:          input.Amount,
		TransactionDate: time.Time(input.TransactionDate),
	}

	created, err := h.PurchaseService.CreatePurchase(purchase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.Error{Error: "internal error"})
		return
	}

	c.JSON(http.StatusOK, schemas.CreatePurchaseOutput{
		ID:              created.ID,
		Description:     created.Description,
		Amount:          created.Amount,
		TransactionDate: created.TransactionDate,
	})

}

// @Produce      json
// @Param        purchase-id  path  int  true  "purchaseId"
// @Success      200  {object}  GetPurchaseOutput
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /purchases/{purchase-id} [get]
func (h *PurchaseHandler) GetPurchase(c *gin.Context) {
	input := schemas.GetPurchaseInput{}

	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, schemas.Error{Error: err.Error()})
		return
	}

	purchase, err := h.PurchaseService.GetPurchaseInCurrency(input.ID, input.Currency)
	if err != nil {
		if err == business.ErrCantConvertCurrency {
			c.JSON(http.StatusInternalServerError, schemas.Error{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, schemas.Error{Error: "internal error"})
		return
	}

	if purchase == nil {
		c.JSON(http.StatusNotFound, schemas.Error{Error: "not found"})
		return
	}

	c.JSON(http.StatusOK, schemas.GetPurchaseOutput{
		ID:              purchase.ID,
		Description:     purchase.Description,
		Amount:          purchase.Amount,
		TransactionDate: purchase.TransactionDate,
		CurrencyRate:    purchase.CurrencyRate,
		ConvertedAmount: purchase.ConvertedAmount,
	})
}
