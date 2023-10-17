package routes

import (
	"go-inventory/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoadRoutes(router *gin.Engine, db *mongo.Database) {
	router.GET("/items", controllers.GetItems(db))
}
