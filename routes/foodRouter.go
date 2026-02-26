package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/food", controller.GETFoods())
	incomingRoutes.GET("/food/:food_id", controller.GETFood())
	incomingRoutes.POST("/food", controller.POSTFood())
	incomingRoutes.PATCH("/food/:food_id", controller.UpdateFood())
}
