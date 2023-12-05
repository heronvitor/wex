package rest

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
	router         *gin.Engine
	AccountHandler PurchaseHandler
}

func (api *API) SetupRouter() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	r.Group("").
		POST("/purchase", api.AccountHandler.CreatePurchase).
		GET("/purchase", api.AccountHandler.GetPurchase)
	api.router = r
}

func (api *API) Run(addr string) (close func() error) {
	srv := &http.Server{
		Addr:    addr,
		Handler: api.router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("listen: %s\n", err)
		}
	}()
	return srv.Close
}
