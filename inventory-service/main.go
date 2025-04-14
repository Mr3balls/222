package main

import (
	"context"
	"log"
	"time"

	"inventory-service/internal/delivery/http"
	"inventory-service/internal/repository"
	"inventory-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	db := client.Database("ecommerce")
	productRepo := repository.NewProductRepository(db)
	productUseCase := usecase.NewProductUseCase(productRepo)

	r := gin.Default()
	http.NewProductHandler(r, productUseCase)

	log.Println("Inventory service running on :8081")
	r.Run(":8081")
}
