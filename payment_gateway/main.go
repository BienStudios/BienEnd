package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var version string = "v26.00.0000_ALPHA"

func main() {
	r := gin.New()

	// Middlewares
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	api := r.Group("/payment/v" + version)
	{
		api.POST("/payments", CreatePaymentHandler)
		api.POST("/webhooks", WebHookHandler)
	}

	srv := &http.Server{
		Addr:         ":8082",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	srv.ListenAndServe()

}

func CreatePaymentHandler(c *gin.Context) {
	// Procesar pago
	c.JSON(http.StatusOK, gin.H{"status": "approved", "id": "pay_12345"})
}

func WebHookHandler(c *gin.Context) {
	// Recibir confirmación del proveedor
	c.JSON(http.StatusOK, gin.H{"received": true})
}
