package handler

import (
	"api-gateway/internal/config"
	"api-gateway/internal/middleware"
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API Gateway is up and running!"})
}

func Login(c *gin.Context) {
	var loginRequest struct {
		UserID string `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := middleware.GenerateJWT(loginRequest.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// =============== ORDER SERVICE =================

func CreateOrder(c *gin.Context) {
	userClaims, ok := authorize(c)
	if !ok {
		return
	}

	forwardRequest(c, config.OrderServiceURL+"/orders", http.MethodPost, userClaims.UserID)
}

func GetOrder(c *gin.Context) {
	userClaims, ok := authorize(c)
	if !ok {
		return
	}

	orderID := c.Param("id")
	forwardRequest(c, config.OrderServiceURL+"/orders/"+orderID, http.MethodGet, userClaims.UserID)
}

func ListOrders(c *gin.Context) {
	userClaims, ok := authorize(c)
	if !ok {
		return
	}

	forwardRequest(c, config.OrderServiceURL+"/orders", http.MethodGet, userClaims.UserID)
}

func UpdateOrder(c *gin.Context) {
	userClaims, ok := authorize(c)
	if !ok {
		return
	}

	orderID := c.Param("id")
	forwardRequest(c, config.OrderServiceURL+"/orders/"+orderID, http.MethodPatch, userClaims.UserID)
}

// =============== INVENTORY SERVICE =================

func CreateProduct(c *gin.Context) {
	_, ok := authorize(c)
	if !ok {
		return
	}

	forwardRequest(c, config.InventoryServiceURL+"/products", http.MethodPost, "")
}

func GetProduct(c *gin.Context) {
	_, ok := authorize(c)
	if !ok {
		return
	}

	id := c.Param("id")
	forwardRequest(c, config.InventoryServiceURL+"/products/"+id, http.MethodGet, "")
}

func ListProducts(c *gin.Context) {
	_, ok := authorize(c)
	if !ok {
		return
	}

	forwardRequest(c, config.InventoryServiceURL+"/products", http.MethodGet, "")
}

func UpdateProduct(c *gin.Context) {
	_, ok := authorize(c)
	if !ok {
		return
	}

	id := c.Param("id")
	forwardRequest(c, config.InventoryServiceURL+"/products/"+id, http.MethodPatch, "")
}

func DeleteProduct(c *gin.Context) {
	_, ok := authorize(c)
	if !ok {
		return
	}

	id := c.Param("id")
	forwardRequest(c, config.InventoryServiceURL+"/products/"+id, http.MethodDelete, "")
}

// =============== HELPERS =================

func forwardRequest(c *gin.Context, url, method, userID string) {
	client := &http.Client{}

	var body io.Reader
	if method == http.MethodPost || method == http.MethodPatch {
		jsonBody, _ := io.ReadAll(c.Request.Body)
		body = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Request build error"})
		return
	}
	req.Header = c.Request.Header
	if userID != "" {
		req.Header.Set("X-User-ID", userID)
	}

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service unavailable"})
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
}

func authorize(c *gin.Context) (*middleware.Claims, bool) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return nil, false
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	claims, err := middleware.ParseJWT(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return nil, false
	}
	return claims, true
}
