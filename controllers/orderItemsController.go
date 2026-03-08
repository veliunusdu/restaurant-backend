package controllers

import (
"context"
"golang-restaurant-management/database"
"net/http"
"time"

"github.com/gin-gonic/gin"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/bson/primitive"
"go.mongodb.org/mongo-driver/mongo"
)

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

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

func ItemsByOrder(id string) (orderItems []primitive.M, err error) {
ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
defer cancel()

matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "order_id", Value: id}}}}
lookupStage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "food"}, {Key: "localField", Value: "food_id"}, {Key: "foreignField", Value: "food_id"}, {Key: "as", Value: "food"}}}}
unwindStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$food"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}

lookupOrderStage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "order"}, {Key: "localField", Value: "order_id"}, {Key: "foreignField", Value: "order_id"}, {Key: "as", Value: "order"}}}}
unwindStage2 := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$order"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}
lookupTableStage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "table"}, {Key: "localField", Value: "order.table_id"}, {Key: "foreignField", Value: "table_id"}, {Key: "as", Value: "table"}}}}
unwindTableStage := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$table"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}}
projectStage := bson.D{{Key: "$project", Value: bson.D{
{Key: "_id", Value: 0},
{Key: "amount", Value: "$food.price"},
{Key: "total_count", Value: 1},
{Key: "food_name", Value: "$food.name"},
{Key: "food_image", Value: "$food.image"},
{Key: "table_number", Value: "$table.table_number"},
{Key: "table_id", Value: "$table.table_id"},
{Key: "order_id", Value: "$order.order_id"},
{Key: "price", Value: "$food.price"},
{Key: "quantity", Value: 1},
}}}

groupStage := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: bson.D{{Key: "order_id", Value: "$order_id"}, {Key: "table_id", Value: "$table_id"}, {Key: "table_number", Value: "$table_number"}}}, {Key: "total_amount", Value: bson.D{{Key: "$sum", Value: "$amount"}}}, {Key: "total_count", Value: bson.D{{Key: "$sum", Value: "$total_count"}}}, {Key: "items", Value: bson.D{{Key: "$push", Value: bson.D{{Key: "food_name", Value: "$food_name"}, {Key: "food_image", Value: "$food_image"}, {Key: "price", Value: "$price"}, {Key: "quantity", Value: "$quantity"}}}}}}}}

projectStage2 := bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 0}, {Key: "order_id", Value: "$_id.order_id"}, {Key: "table_id", Value: "$_id.table_id"}, {Key: "table_number", Value: "$_id.table_number"}, {Key: "total_amount", Value: 1}, {Key: "total_count", Value: 1}, {Key: "items", Value: 1}}}}

result, err := orderItemCollection.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage, unwindStage, lookupOrderStage, unwindStage2, lookupTableStage, unwindTableStage, projectStage, groupStage, projectStage2})

if err != nil {
return nil, err
}

if err = result.All(ctx, &orderItems); err != nil {
return nil, err
}

return orderItems, err
}

func UpdateOrderItem() gin.HandlerFunc {
return func(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{"message": "UpdateOrderItem stub"})
}
}
