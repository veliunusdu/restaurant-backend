package controllers

import (
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", GetInvoices())
	incomingRoutes.GET("/invoice/:invoice_id", GetInvoice())
	incomingRoutes.POST("/invoice", CreateInvoice())
	incomingRoutes.PATCH("/invoice/:invoice_id", UpdateInvoice())
}

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
