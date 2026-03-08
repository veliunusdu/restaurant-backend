package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "GetOrderItems stub"})
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "GetOrderItem stub"})
	}
}

func GetOrderItemsByOrderID() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "GetOrderItemsByOrderID stub"})
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{"message": "CreateOrderItem stub"})
	}
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {
	return nil, nil
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "UpdateOrderItem stub"})
	}
}
