package main

import (
	handler "api-gateway/internal/delivery"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", handler.HealthCheck)

	r.POST("/auth/login", handler.Login)

	// =============== ORDER SERVICE =================
	r.POST("/orders", handler.CreateOrder)
	r.GET("/orders/:id", handler.GetOrder)
	r.GET("/orders", handler.ListOrders)
	r.PATCH("/orders/:id", handler.UpdateOrder)

	// =============== INVENTORY SERVICE =================
	r.POST("/products", handler.CreateProduct)
	r.GET("/products/:id", handler.GetProduct)
	r.GET("/products", handler.ListProducts)
	r.PATCH("/products/:id", handler.UpdateProduct)
	r.DELETE("/products/:id", handler.DeleteProduct)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
