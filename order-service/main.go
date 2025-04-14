package main

import (
	"context"
	"log"
	"time"

	"order-service/internal/delivery/http"
	"order-service/internal/repository"
	"order-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Init MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Connect to MongoDB and wait for the connection
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Ensure disconnection on exit
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	db := client.Database("ecommerce")
	orderRepo := repository.NewOrderRepository(db)
	orderUseCase := usecase.NewOrderUseCase(orderRepo)

	// Init Gin
	r := gin.Default()
	http.NewOrderHandler(r, orderUseCase)

	log.Println("Order service running on :8082")
	r.Run(":8082")
}
