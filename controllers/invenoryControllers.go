package controllers

import (
	"context"
	"go-inventory/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetItems(db *mongo.Database) func(*gin.Context) {
	return func(c *gin.Context) {
		collection := db.Collection("items")

		var items []models.Item

		cur, err := collection.Find(context.Background(), bson.M{})
		if err != nil {
			c.JSON(2400, gin.H{"error": err.Error()})
			return
		}

		defer cur.Close(context.Background())

		for cur.Next(context.Background()) {
			var item models.Item
			cur.Decode(&item)
			items = append(items, item)
		}

		c.JSON(200, items)

	}
}
