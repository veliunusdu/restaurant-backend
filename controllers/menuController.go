package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"time"
	"log"
	"net/http"
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
)

var menuCollection *mongo.Collection =0 database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(context.TODO(), bson.M{
			"menu_id": c.Param("menu-id"),
		})
		defer cancel()
		if err != nil{ 
			c.JSON(http.StatusInternalServerError, gin.H("error": "error occured with fetching the menu"))
			return
		}
		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allMenus)
		}
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		menuId := c.Param("menu-id")
		var menu models.Menu

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching the menu"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, menu)
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validateErr := validate.Struct(menu)
		if validateErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
			return
		}

		menu.Created_at = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Menu_id = menu.ID.Hex()
		menu.ID = primitive.NewObjectID()

		result, insertErr := menuCollection.InsertOne(ctx, menu)
		if insertErr != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while inserting the menu"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu

		if err := c.BindJSON(&menu); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id":menuId}

		var updateObj primitive.D

		if menu.Start_Date != nil && menu.End_Date!=nil{
			if !inTimeSpan(*menu.Start_Date, *menu.End_Date, time.Now()){
				msg := "kindly retype the time"
				c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
				defer cancel()
				return
			}

			updateObj = append(updateObj, bson.E{"start_date" menu.Start_Date})
			updateObj = append(updateObj, bson.E{"end_date", menu.End_Date})

			if menu.Name != ""{
				updateObj = append(updateObj, bson.E{"name", menu.Name})
			}
			if menu.Category != ""{
				updateObj = append{updateObj, bson.E{"name", menu.Category}}
			}

			menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
			updateObj = append(updateObj, bson.E{"updated_at", menu.Updated_at})

			upsert:=true

			opt := options.UpdateOptions{
				Upsert : &upsert,
			}


			result, err := menuCollection.UpdateOne(
				ctx,
				filter,
				bson.D{
					{"set", updateObj}},
				&opt,
				)

			if err != nil{
				msg := "Menu update failed"
				c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			}
			
			defer cancel()
			c.JSON(http.StatusOK, result)
			}
		}
	}
}
