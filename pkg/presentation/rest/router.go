package rest

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/heronvitor/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type API struct {
	router         *gin.Engine
	AccountHandler PurchaseHandler
}

func (api *API) SetupRouter() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		location := url.URL{Path: "/swagger/index.html"}
		c.Redirect(http.StatusFound, location.RequestURI())
	})

	r.GET("/swagger/*any", func(context *gin.Context) {
		docs.SwaggerInfo.Host = context.Request.Host
		ginSwagger.CustomWrapHandler(&ginSwagger.Config{URL: "/swagger/doc.json"}, swaggerFiles.Handler)(context)
	})

	r.Group("/api/v1").
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
