package main

import (
	"context"
	"go-inventory/routes"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	router := gin.Default()

	clientOptions := options.Client().ApplyURI("mongodb+srv://alice:vonataraxia@cluster0.bhdv254.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	db := client.Database("inventory")

	routes.LoadRoutes(router, db)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
