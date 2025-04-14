package repository

import (
	"context"
	"order-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{
		collection: db.Collection("orders"),
	}
}

// CreateOrder сохраняет новый заказ в базе данных
func (r *OrderRepository) CreateOrder(ctx context.Context, order *models.Order) error {
	_, err := r.collection.InsertOne(ctx, order)
	return err
}

// GetOrderByID получает заказ по ID
func (r *OrderRepository) GetOrderByID(ctx context.Context, id string) (*models.Order, error) {
	var order models.Order
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)
	return &order, err
}

// GetAllOrders получает все заказы
func (r *OrderRepository) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// UpdateOrder обновляет статус заказа
func (r *OrderRepository) UpdateOrder(ctx context.Context, id string, order *models.Order) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": order})
	return err
}
