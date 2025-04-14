package repository

import (
	"context"
	"inventory-service/internal/models"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository interface {
	Create(ctx context.Context, p *models.Product) error
	GetByID(ctx context.Context, id string) (*models.Product, error)
	List(ctx context.Context) ([]*models.Product, error)
	Update(ctx context.Context, id string, p *models.Product) error
	Delete(ctx context.Context, id string) error
	ListFiltered(ctx context.Context, category, limit, offset string) ([]*models.Product, error)
}

type productRepo struct {
	col *mongo.Collection
}

func NewProductRepository(db *mongo.Database) ProductRepository {
	return &productRepo{col: db.Collection("products")}
}

func (r *productRepo) Create(ctx context.Context, p *models.Product) error {
	_, err := r.col.InsertOne(ctx, p)
	return err
}

func (r *productRepo) GetByID(ctx context.Context, id string) (*models.Product, error) {
	var p models.Product
	err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&p)
	return &p, err
}

func (r *productRepo) List(ctx context.Context) ([]*models.Product, error) {
	cursor, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*models.Product
	for cursor.Next(ctx) {
		var p models.Product
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	return products, nil
}

func (r *productRepo) Update(ctx context.Context, id string, p *models.Product) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": p}
	_, err := r.col.UpdateOne(ctx, filter, update)
	return err
}

func (r *productRepo) Delete(ctx context.Context, id string) error {
	_, err := r.col.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *productRepo) ListFiltered(ctx context.Context, category, limitStr, offsetStr string) ([]*models.Product, error) {
	filter := bson.M{}
	if category != "" {
		filter["category"] = category
	}

	opts := options.Find()
	if limit, err := strconv.Atoi(limitStr); err == nil {
		opts.SetLimit(int64(limit))
	}
	if offset, err := strconv.Atoi(offsetStr); err == nil {
		opts.SetSkip(int64(offset))
	}

	cursor, err := r.col.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*models.Product
	for cursor.Next(ctx) {
		var p models.Product
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	return products, nil
}
