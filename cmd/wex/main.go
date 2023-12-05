package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alecthomas/kong"
	"github.com/heronvitor/pkg/business"
	"github.com/heronvitor/pkg/dataaccess/clients/fiscaldata"
	"github.com/heronvitor/pkg/dataaccess/repository"
	"github.com/heronvitor/pkg/presentation/rest"
)

func main() {
	config := config{}

	ctx := kong.Parse(&config)
	if ctx.Error != nil {
		log.Fatal(ctx.Error)
	}

	switch ctx.Command() {
	case "update":
		update(config)
	case "api":
		runApi(config)
	default:
		panic(ctx.Command())
	}
}

func runApi(config config) {
	db, err := sql.Open("postgres", config.DbURL)
	if err != nil {
		return
	}

	app := &rest.API{
		AccountHandler: rest.PurchaseHandler{
			PurchaseService: business.TransactionService{
				TransactionRepository: repository.PurchaseRepository{
					DB: db,
				},
				ExchangeRateRepository: repository.ExangeRateRepository{
					DB: db,
				},
			},
		},
	}

	app.SetupRouter()

	app.Run(fmt.Sprintf(":%d", config.API.Port))
	<-context.Background().Done()
	return
}

func update(config config) (err error) {
	db, err := sql.Open("postgres", config.DbURL)
	if err != nil {
		return
	}

	exchangeRatesService := business.ExchangeRatesService{
		ExchangeRateRepository: repository.ExangeRateRepository{
			DB: db,
		},
		ExchangeRatesClient: &fiscaldata.Client{
			Client: &http.Client{Timeout: 10 * time.Second},
			URL:    "https://api.fiscaldata.treasury.gov",
		},
	}

	exchangeRatesService.Update(business.UpdateOptions{
		Now:           time.Now(),
		RetryInterval: time.Hour,
		Interval:      24 * time.Hour,
	})
	return
}
